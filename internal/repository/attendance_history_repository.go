package repository

import (
	"context"

	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AttendanceHistoryRepository interface {
	Create(ctx context.Context, entity *entity.AttendanceHistory) error
	GetMaxClockTimeByEmployeeID(employeeID uuid.UUID) (*model.MaxClockTime, error)
	GetFilteredAttendanceHistories(ctx context.Context, filters *model.AttendanceHistoryFilter) ([]entity.AttendanceHistory, error)
}

type attendanceHistoryRepository struct {
	Repository[entity.AttendanceHistory]
}

func NewAttendanceHistoryRepository(log *logrus.Logger, db *gorm.DB) AttendanceHistoryRepository {
	return &attendanceHistoryRepository{
		Repository: Repository[entity.AttendanceHistory]{
			log,
			db,
		},
	}
}

func (r *attendanceHistoryRepository) GetMaxClockTimeByEmployeeID(employeeID uuid.UUID) (*model.MaxClockTime, error) {
	var result model.MaxClockTime

	if err := r.DB.Table("employees").
		Select("departements.max_clock_in_time, departements.max_clock_out_time").
		Joins("join departements on employees.departement_id = departements.id").
		Where("employees.id = ?", employeeID).
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil

}

func (r *attendanceHistoryRepository) GetFilteredAttendanceHistories(ctx context.Context, filters *model.AttendanceHistoryFilter) ([]entity.AttendanceHistory, error) {
	var attendanceHistories []entity.AttendanceHistory

	query := r.DB.
		Preload("Employee").
		Preload("Employee.Departement").
		Preload("Attendance").
		Joins("JOIN attendances ON attendances.id = attendance_histories.attendance_id").
		Joins("JOIN employees ON employees.id = attendances.employee_id").
		Where("attendance_histories.date_attendance BETWEEN ? AND ?", filters.StartDate, filters.EndDate)

	if filters.DepartementID != uuid.Nil {
		query = query.Where("employees.departement_id = ?", filters.DepartementID)
	}

	if err := query.Find(&attendanceHistories).Error; err != nil {
		return nil, err
	}

	return attendanceHistories, nil

}
