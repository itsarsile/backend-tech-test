package handlers

import (
	"backend/models"
	"backend/pkgs/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ClockIn handles clocking in for an employee
// @Summary Clock in
// @Description Clocks in an employee
// @Tags Attendance
// @Accept json
// @Produce json
// @Param request body models.ClockInRequest true "Clock-in request"
// @Router /clockin [post]
func ClockIn(c *gin.Context) {
	var req models.ClockInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendance := models.Attendance{
		AttendanceID: uuid.New(),
		EmployeeID:   req.EmployeeID,
		ClockIn:      func() *time.Time { t := time.Now(); return &t }(),
		ClockOut:     nil,
	}

	if err := database.DB.Create(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	attendanceHistory := models.AttendanceHistory{
		EmployeeID:     req.EmployeeID,
		AttendanceID:   attendance.AttendanceID.String(),
		DateAttendance: time.Now(),
		AttendanceType: 1,
		Description:    "Clocked in",
	}

	if err := database.DB.Create(&attendanceHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance created successfully", "employee": attendance})

}

// ClockOut handles clocking out for an employee
// @Summary Clock out
// @Description Clocks out an employee
// @Tags Attendance
// @Accept json
// @Produce json
// @Param request body models.ClockInRequest true "Clock-out request"
// @Router /clockout [put]
func ClockOut(c *gin.Context) {
	var req models.ClockInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var attendance models.Attendance
	if err := database.DB.Where("employee_id = ? AND clock_out IS NULL", req.EmployeeID).First(&attendance).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No clock-in found yet for this employee"})
		return
	}

	clockOutTime := time.Now()
	attendance.ClockOut = &clockOutTime

	if err := database.DB.Save(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	attendanceHistory := models.AttendanceHistory{
		EmployeeID:     req.EmployeeID,
		AttendanceID:   attendance.AttendanceID.String(),
		DateAttendance: clockOutTime,
		AttendanceType: 0,
		Description:    "Clocked out",
	}

	if err := database.DB.Create(&attendanceHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance created successfully", "employee": attendance})

}
