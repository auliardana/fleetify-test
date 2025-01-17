package repository

import (
	"context"

	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(ctx context.Context, entity *entity.Employee) error
	FindById(ctx context.Context, id any) (*entity.Employee, error)
	Update(ctx context.Context, entity *entity.Employee) error
	Delete(ctx context.Context, id any) error

	// GetAll retrieves all records from the database.
	GetAll(ctx context.Context) ([]entity.Employee, error)
}

type employeeRepository struct {
	Repository[entity.Employee]
}

func NewEmployeeRepository(log *logrus.Logger, db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		Repository: Repository[entity.Employee]{
			log,
			db,
		},
	}
}

func (r *employeeRepository) GetAll(ctx context.Context) ([]entity.Employee, error) {
	var employees []entity.Employee
	if err := r.DB.WithContext(ctx).Preload("Departement").Find(&employees).Error; err != nil {
		r.Log.Errorf("failed to get all employees: %v", err)
		return nil, err
	}
	return employees, nil
}
