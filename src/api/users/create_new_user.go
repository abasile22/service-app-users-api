package users

import (
	"github.com/abasile22/service-app-api/src/api/driver"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct{
	User string `json:"user"`
	Password string `json:"password"`
}

func CreateNewUser(c *gin.Context) {
	var login Login

	err := c.BindJSON(&login)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
	} else {
		driver.StartMySQL(login.User, login.Password)
		c.JSON(200, gin.H{"status": login.User})
	}

}

