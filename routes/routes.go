package routes

import (
	"employee-task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/tasks", controllers.CreateTask)

	router.GET("/tasks", controllers.GetTasks)

	router.DELETE("/tasks/:id", controllers.DeleteTask)

	router.PUT("/tasks/:id", controllers.UpdateTask)
}
