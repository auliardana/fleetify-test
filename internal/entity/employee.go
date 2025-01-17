package entity

import (
	"time"
)

type Employee struct {
    ID            string      `gorm:"primaryKey;type:varchar(50)" json:"id"`
    DepartementID int         `gorm:"type:int" json:"departement_id"`
    Name          string      `gorm:"type:varchar(255);not null" json:"name"`
    Address       string      `gorm:"type:text" json:"address"`
    CreatedAt     time.Time   `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt     time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
    Departement   Departement `gorm:"foreignKey:DepartementID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"departement"`
}
