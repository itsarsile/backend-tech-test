package models

import (
	"time"
)

type AttendanceHistory struct {
	EmployeeID     string
	AttendanceID   string
	DateAttendance time.Time `gorm:"type:timestamp"`
	AttendanceType int8      `gorm:"type:tinyint(1)"`
	Description    string    `gorm:"type:text"`
}
