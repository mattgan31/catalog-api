package controllers

import (
	"net/http"
	"time"

	"example.com/m/v2/database"
	"example.com/m/v2/helpers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.Created_At = time.Now()
	User.Updated_At = time.Now()

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.Full_Name,
	})
}
