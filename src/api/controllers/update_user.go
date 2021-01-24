package controllers

import (
	"github.com/abasile22/service-app-api/src/api/payload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUser(c *gin.Context) {
	var login payload.Login

	err := c.BindJSON(&login)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
	} else {
		c.JSON(200, gin.H{"status": login.User})
	}

}

