package main

import (
	"backend/models"
	"backend/pkgs/database"
	"backend/routes"

	docs "backend/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.DatabaseConnection()

	database.DB.AutoMigrate(
		&models.Department{},
		&models.Employee{},
		&models.Attendance{},
		&models.AttendanceHistory{},
	)

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	routes.SetupRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8000")
}
