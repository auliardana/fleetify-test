package model

type ClockInAttendanceRequest struct {
	EmployeeID string `json:"employee_id" validate:"required"`
	// ClockIn    time.Time `json:"clock_in" validate:"required"`
}

type ClockOutAttendanceRequest struct {
	EmployeeID string `json:"employee_id" validate:"required"`
	// ClockOut   time.Time `json:"clock_out" validate:"required"`
}
