package handlers

import (
	"backend/models"
	"backend/pkgs/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateDepartment creates a new department
// @Summary Create a new department
// @Description Creates a new department with the provided details
// @Tags Departments
// @Accept json
// @Produce json
// @Param department body models.CreateDepartmentRequest true "Department details"
// @Router /departments [post]
func CreateDepartment(c *gin.Context) {
	var req models.CreateDepartmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	department := models.Department{
		DepartmentName:  req.DepartmentName,
		MaxClockInTime:  req.MaxClockInTime,
		MaxClockOutTime: req.MaxClockOutTime,
	}

	if err := database.DB.Create(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created department
	c.JSON(http.StatusOK, department)
}

// GetAllDepartments retrieves all departments
// @Summary Get all departments
// @Description Retrieves all departments
// @Tags Departments
// @Accept json
// @Produce json
// @Router /departments [get]
func GetAllDepartments(c *gin.Context) {

	department := []models.Department{}

	database.DB.Find(&department)

	c.JSON(http.StatusOK, department)

}

// GetDepartmentById retrieves a department by ID
// @Summary Get a department by ID
// @Description Retrieves a department by the provided ID
// @Tags Departments
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Router /departments/{id} [get]
func GetDepartmentById(c *gin.Context) {

	id := c.Param("id")
	department := []models.Department{}

	database.DB.Find(&department, id)

	c.JSON(http.StatusOK, department)

}

// UpdateDepartment updates a department by ID
// @Summary Update a department by ID
// @Description Updates a department with the provided details
// @Tags Departments
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Param department body models.UpdateDepartmentRequest true "Department details"
// @Router /departments/{id} [put]
func UpdateDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	if err := database.DB.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	var req models.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department.DepartmentName = req.DepartmentName
	department.MaxClockInTime = req.MaxClockInTime
	department.MaxClockOutTime = req.MaxClockOutTime

	if err := database.DB.Save(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, department)

}

// DeleteDepartment deletes a department by ID
// @Summary Delete a department by ID
// @Description Deletes a department with the provided ID
// @Tags Departments
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Router /departments/{id} [delete]
func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	if err := database.DB.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	if err := database.DB.Where("id = ?", id).Delete(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, department)

}
