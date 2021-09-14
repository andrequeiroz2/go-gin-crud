package controllers

import (
	
	"gincrud/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

type UpdateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task string `json:"task"`
	Deadline string `json:"deadline"`
}

func FindTasks(c *gin.Context){
	db:= c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {

	var input CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	task := models.Task{
		AssingedTo: input.AssignedTo, 
		Task: input.Task,
		Deadline: deadline,
	}

	db := c.MustGet("db").(*gorm.DB)

	if db.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": &db.Error, "func":"MustGet"})
		return	
	}

	db.Create(&task)

	if db.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": &db.Error, "func":"Create"})
		return	
	}

	c.JSON(http.StatusOK, gin.H{"data": &task})
}


func FindTask(c *gin.Context){

	var task models.Task

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id= ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Record not found"})
		return	
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}


func UpdateTask(c *gin.Context){

	db:= c.MustGet("db").(*gorm.DB)

	var task models.Task

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil{ 
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	date:= "2006-01-02"
	deadline, _:= time.Parse(date, input.Deadline)

	var updatedInput models.Task

	updatedInput.Deadline = deadline
	updatedInput.AssingedTo =  input.AssignedTo
	updatedInput.Task = input.Task

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}


func DeleteTask(c *gin.Context){

	db:= c.MustGet("db").(*gorm.DB)
	var book models.Task

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"erroe": "Record not found"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})

}

