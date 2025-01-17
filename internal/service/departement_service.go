package service

import (
	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/auliardana/fleetify-test/internal/model"
	"github.com/auliardana/fleetify-test/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DepartementService interface {
	CreateDepartement(c *gin.Context, req *model.DepartementRequest) error
	ListDepartement(c *gin.Context) ([]entity.Departement, error)
	// GetDepartementByID(c *gin.Context, req string) (*entity.Departement, error)
	UpdateDepartement(c *gin.Context, req *model.DepartementUpdateRequest) error
	DeleteDepartement(c *gin.Context, id int) error
}

type departementService struct {
	Repo repository.DepartementRepository
	Log  *logrus.Logger
}

func NewDepartementService(repo repository.DepartementRepository, log *logrus.Logger) DepartementService {
	return &departementService{
		Repo: repo,
		Log:  log,
	}
}

func (s *departementService) CreateDepartement(c *gin.Context, req *model.DepartementRequest) error {

	departement := &entity.Departement{
		DepartementName: req.DepartementName,
		MaxClockInTime:  req.MaxClockInTime,
		MaxClockOutTime: req.MaxClockOutTime,
	}

	err := s.Repo.Create(c, departement)
	if err != nil {
		s.Log.Warn("Failed to create departement: ", err)
		return err
	}

	return nil

}

func (s *departementService) ListDepartement(c *gin.Context) ([]entity.Departement, error) {
	departements, err := s.Repo.GetAll(c)
	if err != nil {
		s.Log.Warn("Failed to list departement: ", err)
		return nil, err
	}
	return departements, nil

}

// func (s *departementService) GetDepartementByID(c *gin.Context, id string) (*entity.Departement, error) {

// }

func (s *departementService) UpdateDepartement(c *gin.Context, req *model.DepartementUpdateRequest) error {
	//find departement by id
	data , err := s.Repo.FindById(c, req.ID)
	if err != nil {
		s.Log.Warn("Failed to get departement: ", err)
		return err
	}

	if req.DepartementName != "" {
		data.DepartementName = req.DepartementName
	}

	if !req.MaxClockInTime.IsZero() {
		data.MaxClockInTime = req.MaxClockInTime
	}

	if !req.MaxClockOutTime.IsZero() {
		data.MaxClockOutTime = req.MaxClockOutTime
	}

	err = s.Repo.Update(c, data)
	if err != nil {
		s.Log.Warn("Failed to update departement: ", err)
		return err
	}

	return nil

}

func (s *departementService) DeleteDepartement(c *gin.Context, id int) error {
	
	err := s.Repo.Delete(c, id)
	if err != nil {
		s.Log.Warn("Failed to delete departement: ", err)
		return err
	}

	return nil
}
