package model

import "time"

type AttendanceHistoryFilter struct {
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	DepartementID  string    `json:"departement_id"`
	AttendanceType int       `json:"attendance_type"`
}

type AttendanceHistoryResponse struct {
	ID              string    `json:"id"`
	EmployeeID      string    `json:"employee_id"`
	EmployeeName    string    `json:"employee_name"`
	Departement     string    `json:"departement"`
	DateAttendance  time.Time `json:"date_attendance"`
	AttendanceType  string    `json:"attendance_type"`
	Description     string    `json:"description"`
	ClockIn         time.Time `json:"clock_in"`
	ClockOut        time.Time `json:"clock_out"`
	MaxClockInTime  time.Time `json:"max_clock_in_time"`
	MaxClockOutTime time.Time `json:"max_clock_out_time"`
}
