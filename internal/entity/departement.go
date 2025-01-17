package entity

import (
	"time"
)

type Departement struct {
	ID              int       `gorm:"primaryKey;autoIncrement" json:"id"`
	DepartementName string    `gorm:"type:varchar(255)" json:"departement_name"`
	MaxClockInTime  time.Time `gorm:"column:max_clock_in_time" json:"max_clock_in_time"`
	MaxClockOutTime time.Time `gorm:"column:max_clock_out_time" json:"max_clock_out_time"`
}
