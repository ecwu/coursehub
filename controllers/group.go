package controllers

import (
	"coursehub/db"
	"coursehub/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGroup(c *gin.Context) {
	groupName := c.PostForm("group_name")
	subName := c.PostForm("sub_name")
	abbrName := c.PostForm("abbr_name")
	subGroupTo := c.PostForm("sub_group_to")
	group := models.GroupModel{
		GroupName: groupName,
		SubName:   subName,
		AbbrName:  abbrName,
	}
	db.GetDB().Save(&group)
	if subGroupTo != "" {
		var parentGroup models.GroupModel
		db.GetDB().Find(&parentGroup, subGroupTo)
		db.GetDB().Model(&group).Update("group_model_id", parentGroup.ID)
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Group created successfully",
		"resourceId": group.ID,
	})
}

func FetchAllGroup(c *gin.Context) {
	var groups []models.GroupModel
	db.GetDB().Preload("SubGroupTo").Find(&groups)
	if len(groups) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No group found."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": groups})
}
