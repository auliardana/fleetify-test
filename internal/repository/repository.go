package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	Log *logrus.Logger
	DB  *gorm.DB
}

func NewRepository[T any](log *logrus.Logger, db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		Log: log,
		DB:  db,
	}
}

// Create inserts a new record into the database.
func (r *Repository[T]) Create(ctx context.Context, entity *T) error {
	if err := r.DB.WithContext(ctx).Create(entity).Error; err != nil {
		r.Log.Errorf("failed to create entity: %v", err)
		return err
	}
	return nil
}

// Read retrieves a record by its primary key.
func (r *Repository[T]) FindById(ctx context.Context, id any) (*T, error) {
	var entity T
	if err := r.DB.WithContext(ctx).First(&entity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Log.Warnf("record not found: %v", err)
			return nil, nil
		}
		r.Log.Errorf("failed to read entity: %v", err)
		return nil, err
	}
	return &entity, nil
}

// Update modifies an existing record.
func (r *Repository[T]) Update(ctx context.Context, entity *T) error {
	result := r.DB.WithContext(ctx).Save(entity)
	if result.Error != nil {
		r.Log.Errorf("failed to update entity: %v", result.Error)
		return result.Error
	}

	return nil
}

// Delete removes a record by its primary key.
func (r *Repository[T]) Delete(ctx context.Context, id any, ) error {
	result := r.DB.WithContext(ctx).Delete(new(T), "id = ?", id)
	if result.Error != nil {
		r.Log.Errorf("failed to delete entity: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
