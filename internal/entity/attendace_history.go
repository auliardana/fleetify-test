package entity

import "time"

type AttendanceHistory struct {
	ID           string     `gorm:"primaryKey;autoIncrement"`
	EmployeeID   string     `gorm:"type:varchar(50);not null"`
	AttendanceID string     `gorm:"type:varchar(100);not null"`
	Employee     Employee   `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attendance   Attendance `gorm:"foreignKey:AttendanceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DateAttendance time.Time `gorm:"type:timestamp"`
	AttendanceType uint8     `gorm:"type:tinyint(1)"`
	Description    string    `gorm:"type:text"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
