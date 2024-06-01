package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	EmployeeID   string `gorm:"type:varchar(50);unique"`
	Name         string `gorm:"type:varchar(255)"`
	Address      string `gorm:"type:text"`
	DepartmentID uint

	Attendances         []Attendance        `gorm:"foreignKey:EmployeeID;references:EmployeeID"`
	AttendanceHistories []AttendanceHistory `gorm:"foreignKey:EmployeeID;references:EmployeeID"`
}

type CreateEmployeeRequest struct {
	EmployeeID   string `json:"employee_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	DepartmentID uint   `json:"department_id" binding:"required"`
}
