package models

import (
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey"`
	DepartmentName  string `gorm:"type:varchar(255)"`
	MaxClockInTime  string `gorm:"type:time(3)"`
	MaxClockOutTime string `gorm:"type:time(3)"`

	Employees []Employee `gorm:"foreignKey:DepartmentID"`
}

type CreateDepartmentRequest struct {
	DepartmentName  string `json:"department_name" binding:"required"`
	MaxClockInTime  string `json:"max_clock_in_time" binding:"required"`
	MaxClockOutTime string `json:"max_clock_out_time" binding:"required"`
}

type UpdateDepartmentRequest struct {
	DepartmentName  string `json:"department_name"`
	MaxClockInTime  string `json:"max_clock_in_time"`
	MaxClockOutTime string `json:"max_clock_out_time"`
}
