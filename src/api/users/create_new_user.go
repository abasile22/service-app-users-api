package users

import (
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
		c.JSON(200, gin.H{"status": login.User})
	}

}

