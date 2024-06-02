package routes

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	departments := api.Group("/departments")
	{
		departments.GET("", handlers.GetAllDepartments)
		departments.GET("/:id", handlers.GetDepartmentById)
		departments.POST("", handlers.CreateDepartment)
		departments.PUT("/:id", handlers.UpdateDepartment)
		departments.DELETE("/:id", handlers.DeleteDepartment)
	}

	employees := api.Group("/employees")
	{
		employees.POST("", handlers.CreateEmployee)
		employees.GET("", handlers.GetAllEmployee)
		employees.GET("/:id", handlers.GetEmployeeById)
		employees.PUT("/:id", handlers.UpdateEmployee)
		employees.DELETE("/:id", handlers.DeleteEmployee)
	}

	api.POST("/clockin", handlers.ClockIn)
	api.PUT("/clockout", handlers.ClockOut)

	attendances := api.Group("/attendances")
	{
		attendances.GET("/histories", handlers.GetAttendanceHistories)
		attendances.GET("/log", handlers.GetAttendanceLog)
	}
}
