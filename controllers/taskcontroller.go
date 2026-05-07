package controllers

import (
	"employee-task-manager/database"
	"employee-task-manager/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var task models.Task

	c.BindJSON(&task)

	database.DB.Create(&task)

	c.JSON(200, gin.H{
		"message": "Task Created",
		"data":    task,
	})
}

func GetTasks(c *gin.Context) {

	var tasks []models.Task

	database.DB.Find(&tasks)

	c.JSON(200, gin.H{
		"data": tasks,
	})
}
func DeleteTask(c *gin.Context) {

	id := c.Param("id")

	var task models.Task

	database.DB.First(&task, id)

	database.DB.Delete(&task)

	c.JSON(200, gin.H{

		"message": "Task Deleted",
	})
}
func UpdateTask(c *gin.Context) {

	id := c.Param("id")

	var task models.Task

	database.DB.First(&task, id)

	c.BindJSON(&task)

	database.DB.Save(&task)

	c.JSON(200, gin.H{

		"message": "Task Updated",

		"data": task,
	})
}
