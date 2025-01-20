package model

import "github.com/google/uuid"

type EmployeeRequest struct {
	DepartementID uuid.UUID `json:"departement_id" validate:"required"`
	Name          string    `json:"name" validate:"required,max=255"`
	Address       string    `json:"address" gorm:"type:text"`
}

type EmployeeUpdateRequest struct {
	ID            uuid.UUID `json:"id"`
	DepartementID uuid.UUID `json:"departement_id"`
	Name          string    `json:"name" validate:"max=255"`
	Address       string    `json:"address" gorm:"type:text"`
}
