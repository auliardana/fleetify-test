package repository

import (
	"context"

	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Create(ctx context.Context, entity *entity.Attendance) error
	FindById(ctx context.Context, id any) (*entity.Attendance, error)
	Update(ctx context.Context, entity *entity.Attendance) error
	Delete(ctx context.Context, id any) error

	FindByEmployeeID(ctx context.Context, employeeID string) (*entity.Attendance, error)
}

type attendanceRepository struct {
	Repository[entity.Attendance]
}

func NewAttendanceRepository(log *logrus.Logger, db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{
		Repository: Repository[entity.Attendance]{
			log,
			db,
		},
	}
}

func (r *attendanceRepository) FindByEmployeeID(ctx context.Context, employeeID string) (*entity.Attendance, error) {
	var attendance entity.Attendance
	if err := r.DB.WithContext(ctx).Where("employee_id = ?", employeeID).First(&attendance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Log.Warnf("record not found: %v", err)
			return nil, nil
		}
		r.Log.Errorf("failed to read entity: %v", err)
		return nil, err
	}
	return &attendance, nil

}
