package service

import (
	"fmt"
	"time"

	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/auliardana/fleetify-test/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AttendanceService interface {
	ClockIn(c *gin.Context, req *model.ClockInAttendanceRequest) error
	ClockOut(c *gin.Context, id string) error
}

type attendanceService struct {
	RepoAttendance        repository.AttendanceRepository
	RepoAttendanceHistory repository.AttendanceHistoryRepository
	Log                   *logrus.Logger
}

func NewAttendanceService(repoAttendance repository.AttendanceRepository, repoAttendanceHistory repository.AttendanceHistoryRepository, log *logrus.Logger) AttendanceService {
	return &attendanceService{
		RepoAttendance:        repoAttendance,
		RepoAttendanceHistory: repoAttendanceHistory,
		Log:                   log,
	}
}

func (s *attendanceService) ClockIn(c *gin.Context, req *model.ClockInAttendanceRequest) error {
	attendance := &entity.Attendance{
		ID:         uuid.NewString(),
		EmployeeID: req.EmployeeID,
		ClockIn:    time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := s.RepoAttendance.Create(c, attendance)
	if err != nil {
		s.Log.Warn("Failed to clock in: ", err)
		return err
	}

	return nil
}

func (s *attendanceService) ClockOut(c *gin.Context, id string) error {
	attendance, err := s.RepoAttendance.FindByEmployeeID(c, id)
	if err != nil {
		s.Log.Warn("Failed to find attendance: ", err)
		return err
	}

	attendance.ClockOut = time.Now()
	attendance.UpdatedAt = time.Now()

	err = s.RepoAttendance.Update(c, attendance)
	if err != nil {
		s.Log.Warn("Failed to clock out: ", err)
		return err
	}

	attendanceHistory := &entity.AttendanceHistory{
		ID:             uuid.NewString(),
		EmployeeID:     attendance.EmployeeID,
		AttendanceID:   attendance.ID,
		DateAttendance: time.Now(),
		CreatedAt:      attendance.CreatedAt,
		UpdatedAt:      attendance.UpdatedAt,
	}

	maxClockTime, err := s.RepoAttendanceHistory.GetMaxClockTimeByEmployeeID(attendance.EmployeeID)
	if err != nil {
		s.Log.Warn("Failed to get max clock time: ", err)
		return err
	}

	// Validasi maxClockTime
	if maxClockTime == nil {
		s.Log.Warn("MaxClockTime is null for employee ID: ", attendance.EmployeeID)
		return fmt.Errorf("max clock time not found")
	}

	clockIn := attendance.ClockIn
	clockOut := attendance.ClockOut

	switch {
	default:
		s.Log.Warn("Unrecognized attendance type for employee ID: ", attendance.EmployeeID)
		return fmt.Errorf("unrecognized attendance type")
	// Hadir tepat waktu
	case clockIn.Before(maxClockTime.MaxClockInTime) && clockOut.Before(maxClockTime.MaxClockOutTime):
		attendanceHistory.AttendanceType = 1
		attendanceHistory.Description = "Tepat Waktu"

	// // Hadir dengan pulang cepat (ClockOut lebih awal dari batas waktu - 2 jam)
	// case clockIn.Before(maxClockTime.MaxClockInTime) && clockOut.Before(maxClockTime.MaxClockOutTime.Add(-2*time.Hour)):
	// 	attendanceHistory.AttendanceType = 2

	// Telat hadir (ClockIn lebih dari MaxClockInTime)
	case clockIn.After(maxClockTime.MaxClockInTime):
		attendanceHistory.AttendanceType = 2
		attendanceHistory.Description = "Terlambat"
	}

	err = s.RepoAttendanceHistory.Create(c, attendanceHistory)
	if err != nil {
		s.Log.Warn("Failed to create attendance history: ", err)
		return err
	}

	return nil

}
