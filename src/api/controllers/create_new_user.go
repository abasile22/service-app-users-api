package controllers

import (
	"github.com/abasile22/service-app-api/src/api/dao"
	"github.com/abasile22/service-app-api/src/api/payload"
	"github.com/abasile22/service-app-api/src/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNewUser(c *gin.Context) {

	var login payload.Login

	hash := utils.CreateHash(login.Password)

	err := c.BindJSON(&login)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
	} else {
		dao.StartMySQL(login.User, hash)
		c.JSON(200, gin.H{"status": login.User})
	}

}

