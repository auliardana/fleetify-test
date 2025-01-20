package entity

import (
	"time"

	"github.com/google/uuid"
)

// type Attendance struct {
// 	ID         string    `gorm:"primaryKey;type:varchar(100)"`
// 	EmployeeID string    `gorm:"type:varchar(50);not null"`
// 	ClockIn    time.Time `gorm:"type:timestamp"`
// 	ClockOut   time.Time `gorm:"type:timestamp;default:null"`
// 	CreatedAt  time.Time `gorm:"autoCreateTime"`
// 	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
// 	Employee   Employee  `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// }

type Attendance struct {
	ID         uuid.UUID `gorm:"primaryKey;type:char(36)"`
	EmployeeID uuid.UUID `gorm:"type:char(36);not null"`
	ClockIn    time.Time `gorm:"type:timestamp"`
	ClockOut   time.Time `gorm:"type:timestamp;default:null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	Employee   Employee  `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
