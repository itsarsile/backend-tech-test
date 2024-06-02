package handlers

import (
	"backend/models"
	"backend/pkgs/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateEmployee creates a new employee
// @Summary Create a new employee
// @Description Creates a new employee with the provided details
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body models.CreateEmployeeRequest true "Employee details"
// @Router /employees [post]
func CreateEmployee(c *gin.Context) {
	var req models.CreateEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var department models.Department

	if result := database.DB.Where("id = ? AND deleted_at IS NULL", req.DepartmentID).First(&department); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "department not found"})
		return
	}

	employee := models.Employee{
		Name:         req.Name,
		EmployeeID:   req.EmployeeID,
		Address:      req.Address,
		DepartmentID: req.DepartmentID,
	}

	if err := database.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee created successfully", "employee": employee})
}

// GetAllEmployee retrieves all employees
// @Summary Get all employees
// @Description Retrieves all employees
// @Tags Employees
// @Accept json
// @Produce json
// @Router /employees [get]
func GetAllEmployee(c *gin.Context) {

	employee := []models.EmployeeWithDepartment{}

	query := database.DB.Table("employees").
		Select("employees.id, employees.employee_id, employees.name, employees.address,employees.department_id, departments.department_name").
		Joins("JOIN departments ON employees.department_id = departments.id").
		Where("employees.deleted_at IS NULL").
		Scan(&employee)

	if query.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": query.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)

}

// GetEmployeeById retrieves an employee by ID
// @Summary Get an employee by ID
// @Description Retrieves an employee by the provided ID
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Router /employees/{id} [get]
func GetEmployeeById(c *gin.Context) {

	id := c.Param("id")

	var employee models.EmployeeWithDepartment

	query := database.DB.Table("employees").
		Select("employees.id, employees.employee_id, employees.name, employees.address, employees.department_id, departments.department_name").
		Joins("JOIN departments ON employees.department_id = departments.id").
		Where("employees.id = ?", id).
		Scan(&employee)

	if query.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": query.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)

}

// UpdateEmployee updates an employee by ID
// @Summary Update an employee by ID
// @Description Updates an employee with the provided details
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param employee body models.CreateEmployeeRequest true "Employee details"
// @Router /employees/{id} [put]
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee

	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	var req models.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.EmployeeID = req.EmployeeID
	employee.Name = req.Name
	employee.Address = req.Address
	employee.DepartmentID = req.DepartmentID

	if err := database.DB.Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, employee)

}

// DeleteEmployee deletes an employee by ID
// @Summary Delete an employee by ID
// @Description Deletes an employee with the provided ID
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Router /employees/{id} [delete]
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee

	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	if err := database.DB.Where("id = ?", id).Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, employee)

}
