package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	EmployeeID   string
	ID           uint       `gorm:"primaryKey"`
	AttendanceID uuid.UUID  `gorm:"type:varchar(100);unique"`
	ClockIn      *time.Time `gorm:"type:timestamp"`
	ClockOut     *time.Time `gorm:"type:timestamp"`

	// Ref
	AttendanceHistories []AttendanceHistory `gorm:"foreignKey:AttendanceID;references:AttendanceID"`
}

type ClockInRequest struct {
	EmployeeID string `json:"employee_id" binding:"required"`
}
