package handlers

import (
	"backend/models"
	"backend/pkgs/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttDiffInfo struct {
	DepartmentName string
	EmployeeID     string
	ClockInDiff    string
	ClockOutDiff   string
}

// @Summary Get attendance histories
// @Description Retrieves attendance histories based on department and date filter
// @Tags Attendance
// @Accept json
// @Produce json
// @Param department query string false "Department name"
// @Param date query string false "Date filter (YYYY-MM-DD)"
// @Router /attendances/histories [get]
func GetAttendanceHistories(c *gin.Context) {
	department := c.Query("department")
	dateFilter := c.Query("date")

	fmt.Printf("Query Parameters - Department: %s, Date: %s\n", department, dateFilter)

	var attendanceDiffs []AttDiffInfo

	// Yep raw SQL query FTW :D
	query := database.DB.Table("attendance_histories").
		Select("departments.department_name, attendance_histories.employee_id, " +
			"TIMEDIFF(departments.max_clock_in_time, TIME(MAX(CASE WHEN attendance_type = 1 THEN date_attendance END))) AS clock_in_diff, " +
			"TIMEDIFF(departments.max_clock_out_time, TIME(MAX(CASE WHEN attendance_type = 0 THEN date_attendance END))) AS clock_out_diff").
		Joins("JOIN employees ON employees.employee_id = attendance_histories.employee_id").
		Joins("JOIN departments ON departments.id = employees.department_id")

	if department != "" {
		query = query.Where("departments.department_name = ?", department)
	}
	if dateFilter != "" {
		query = query.Where("DATE(attendance_histories.date_attendance) = ?", dateFilter)
	}

	query = query.Group("departments.department_name, attendance_histories.employee_id")

	if err := query.Scan(&attendanceDiffs).Error; err != nil {
		fmt.Printf("Error fetching attendance differences: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching attendance differences"})
		return
	}

	c.JSON(http.StatusOK, attendanceDiffs)
}

// GetAttendanceLog retrieves all attendance logs
// @Summary Get all attendance logs
// @Description Retrieves all attendance logs
// @Tags Attendance
// @Accept json
// @Produce json
// @Router /attendances/log [get]
func GetAttendanceLog(c *gin.Context) {

	attlog := []models.AttendanceHistory{}

	database.DB.Find(&attlog)

	c.JSON(http.StatusOK, attlog)

}
