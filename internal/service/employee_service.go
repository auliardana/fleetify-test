package service

import (
	"time"

	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/auliardana/fleetify-test/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type EmployeeService interface {
	CreateEmployee(c *gin.Context, req *model.EmployeeRequest) error
	ListEmployee(c *gin.Context) ([]entity.Employee, error)
	UpdateEmployee(c *gin.Context, req *model.EmployeeUpdateRequest) error
	DeleteEmployee(c *gin.Context, id uuid.UUID) error
}

type employeeService struct {
	Repo repository.EmployeeRepository
	Log  *logrus.Logger
}

func NewEmployeeService(repo repository.EmployeeRepository, log *logrus.Logger) EmployeeService {
	return &employeeService{
		Repo: repo,
		Log:  log,
	}
}

func (s *employeeService) CreateEmployee(c *gin.Context, req *model.EmployeeRequest) error {
	employee := &entity.Employee{
		ID:            uuid.New(),
		DepartementID: req.DepartementID,
		Name:          req.Name,
		Address:       req.Address,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := s.Repo.Create(c, employee)
	if err != nil {
		s.Log.Warn("Failed to create employee: ", err)
		return err
	}

	return nil
}

func (s *employeeService) ListEmployee(c *gin.Context) ([]entity.Employee, error) {
	employees, err := s.Repo.GetAll(c)
	if err != nil {
		s.Log.Warn("Failed to list employee: ", err)
		return nil, err
	}

	return employees, nil

}

func (s *employeeService) UpdateEmployee(c *gin.Context, req *model.EmployeeUpdateRequest) error {
	//find departement by id
	data, err := s.Repo.FindById(c, req.ID)
	if err != nil {
		s.Log.Warn("Failed to get departement: ", err)
		return err
	}

	if req.Address != "" {
		data.Address = req.Address
	}

	if req.Name != "" {
		data.Name = req.Name
	}

	if req.DepartementID != uuid.Nil {
		data.DepartementID = req.DepartementID
	}

	err = s.Repo.Update(c, data)
	if err != nil {
		s.Log.Warn("Failed to update departement: ", err)
		return err
	}

	return nil

}

func (s *employeeService) DeleteEmployee(c *gin.Context, id uuid.UUID) error {
	err := s.Repo.Delete(c, id)
	if err != nil {
		s.Log.Warn("Failed to delete employee: ", err)
		return err
	}

	return nil

}
