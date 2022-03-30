package controllers

import (
	"coursehub/db"
	"coursehub/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var err error

func CreateCourse(c *gin.Context) {
	mainTitle := c.PostForm("main_title")
	subTitle := c.PostForm("sub_title")
	courseCode := c.PostForm("course_code")
	courseUnits := uint(c.GetInt("course_units"))
	courseDescription := c.PostForm("course_description")
	courseType := c.PostForm("course_type")
	courseIsVisible := c.GetBool("course_is_visible")
	courseAcceptComment := c.GetBool("course_accept_comment")
	course := models.CourseModel{
		MainTitle:           mainTitle,
		SubTitle:            subTitle,
		CourseCode:          courseCode,
		CourseUnits:         courseUnits,
		CourseDescription:   courseDescription,
		CourseType:          courseType,
		CourseIsVisible:     courseIsVisible,
		CourseAcceptComment: courseAcceptComment,
	}
	db.GetDB().Save(&course)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Course created successfully",
		"resourceId": course.ID,
	})
}

func FetchAllCourse(c *gin.Context) {
	var courses []models.CourseModel
	doUnscopedSearch, err := strconv.ParseBool(c.Request.Header["Unscoped"][0])

	if doUnscopedSearch == true && err == nil {
		db.GetDB().Unscoped().Find(&courses)
	} else {
		db.GetDB().Find(&courses)
	}
	if len(courses) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No course found."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": courses})
}

func FetchSingleCourse(c *gin.Context) {
	var course models.CourseModel
	courseId := c.Param("id")
	db.GetDB().Find(&course, courseId)
	if course.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No course found."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": course})
}

func UpdateCourse(c *gin.Context) {
	var course models.CourseModel
	courseId := c.Param("id")
	db.GetDB().Find(&course, courseId)
	if course.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No course found."})
		return
	}
	mainTitle := c.PostForm("main_title")
	db.GetDB().Model(&course).Update("main_title", mainTitle)
	subTitle := c.PostForm("sub_title")
	db.GetDB().Model(&course).Update("sub_title", subTitle)
	courseCode := c.PostForm("course_code")
	db.GetDB().Model(&course).Update("course_code", courseCode)
	courseUnits := uint(c.GetInt("course_units"))
	db.GetDB().Model(&course).Update("course_units", courseUnits)
	courseDescription := c.PostForm("course_description")
	db.GetDB().Model(&course).Update("course_description", courseDescription)
	courseType := c.PostForm("course_type")
	db.GetDB().Model(&course).Update("course_type", courseType)
	courseIsVisible := c.GetBool("course_is_visible")
	db.GetDB().Model(&course).Update("course_is_visible", courseIsVisible)
	courseAcceptComment := c.GetBool("course_accept_comment")
	db.GetDB().Model(&course).Update("course_accept_comment", courseAcceptComment)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "course update successfully."})
}

func DeleteCourse(c *gin.Context) {
	var course models.CourseModel
	courseId := c.Param("id")
	db.GetDB().Find(&course, courseId)
	if course.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No course found."})
		return
	}
	db.GetDB().Delete(&course)
	c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "course deleted."})
}
