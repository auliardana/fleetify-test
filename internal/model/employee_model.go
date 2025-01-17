package model

type EmployeeRequest struct {
	DepartementID int    `json:"departement_id" validate:"required,max=11"`
	Name          string `json:"name" validate:"required,max=255"`
	Address       string `json:"address" gorm:"type:text"`
}

type EmployeeUpdateRequest struct {
	ID            string `json:"id"`
	DepartementID int    `json:"departement_id" validate:"max=11"`
	Name          string `json:"name" validate:"max=255"`
	Address       string `json:"address" gorm:"type:text"`
}