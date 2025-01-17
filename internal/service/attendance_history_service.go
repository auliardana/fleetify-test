package service

import (
	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/auliardana/fleetify-test/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AttendanceHistoryService interface {
	ListAttendanceHistory(c *gin.Context, filters *model.AttendanceHistoryFilter) ([]entity.AttendanceHistory, error)
}

type attendanceHistoryService struct {
	Repo repository.AttendanceHistoryRepository
	Log  *logrus.Logger
}

func NewAttendanceHistoryService(repo repository.AttendanceHistoryRepository, log *logrus.Logger) AttendanceHistoryService {
	return &attendanceHistoryService{
		Repo: repo,
		Log:  log,
	}
}

func (s *attendanceHistoryService) ListAttendanceHistory(c *gin.Context, filters *model.AttendanceHistoryFilter) ([]entity.AttendanceHistory, error) {
	attendanceHistories, err := s.Repo.GetFilteredAttendanceHistories(c, filters)
	if err != nil {
		s.Log.Warn("Failed to list attendance history: ", err)
		return nil, err
	}

	return attendanceHistories, nil

}