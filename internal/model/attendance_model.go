package model

import "github.com/google/uuid"

type ClockInAttendanceRequest struct {
	EmployeeID uuid.UUID `json:"employee_id" validate:"required"`
	// ClockIn    time.Time `json:"clock_in" validate:"required"`
}

type ClockOutAttendanceRequest struct {
	EmployeeID uuid.UUID `json:"employee_id" validate:"required"`
	// ClockOut   time.Time `json:"clock_out" validate:"required"`
}
