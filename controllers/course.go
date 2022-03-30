package controllers

import (
	"coursehub/db"
	"coursehub/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var err error

func FetchAllCourse(c *gin.Context) {
	var courses []models.CourseModel
	db.GetDB().Find(&courses)
	if len(courses) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No course found."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": courses})
}
