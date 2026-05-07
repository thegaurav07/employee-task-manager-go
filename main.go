package main

import (
	"employee-task-manager/database"
	"employee-task-manager/models"
	"employee-task-manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	database.ConnectDB()

	database.DB.AutoMigrate(&models.Task{})

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	routes.SetupRoutes(router)
	router.GET("/", func(c *gin.Context) {

		c.HTML(200, "index.html", nil)

	})

	router.Run(":8080")
}
