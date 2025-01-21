package model

import (
	"time"

	"github.com/google/uuid"
)

type DepartementRequest struct {
	DepartementName string    `json:"departement_name" validate:"required,max=255"`
	MaxClockInTime  time.Time `json:"max_clock_in_time" validate:"required"`
	MaxClockOutTime time.Time `json:"max_clock_out_time" validate:"required"`
}

type DepartementUpdateRequest struct {
	ID              uuid.UUID `json:"id" validate:"required"`
	DepartementName string    `json:"departement_name" validate:"max=255"`
	MaxClockInTime  time.Time `json:"max_clock_in_time"`
	MaxClockOutTime time.Time `json:"max_clock_out_time"`
}

type MaxClockTime struct {
	MaxClockInTime  time.Time
	MaxClockOutTime time.Time
}
