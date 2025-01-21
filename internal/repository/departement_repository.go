package repository

import (
	"context"

	"github.com/auliardana/fleetify-test/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DepartementRepository interface {
	Create(ctx context.Context, entity *entity.Departement) error
	FindById(ctx context.Context, id any) (*entity.Departement, error)
	Update(ctx context.Context, entity *entity.Departement) error
	Delete(ctx context.Context, id any) error
	
	GetAll(ctx context.Context) ([]entity.Departement, error)
}

type departementRepository struct {
	Repository[entity.Departement]
}

func NewDepartementRepository(log *logrus.Logger, db *gorm.DB) DepartementRepository {
	return &departementRepository{
		Repository: Repository[entity.Departement]{
			log,
			db,
		},
	}
}


func (r *departementRepository) GetAll(ctx context.Context) ([]entity.Departement, error) {
	var departements []entity.Departement
	if err := r.DB.WithContext(ctx).Find(&departements).Error; err != nil {
		r.Log.Errorf("failed to get all departements: %v", err)
		return nil, err
	}
	return departements, nil
}