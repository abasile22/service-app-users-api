package controllers

import (
	"database/sql"
	"fmt"
	"github.com/abasile22/trp-user-role-auth/src/api/payload"
	"github.com/abasile22/trp-user-role-auth/src/api/utils"
	"github.com/gin-gonic/gin"
)

func LogInUser(c *gin.Context) {

	var login payload.Login

	err := c.BindJSON(&login)

	var loginMysql MySql

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/services_app_users")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// See "Important settings" section.
	// perform a db.Query insert
	errr := db.QueryRow("SELECT password FROM users WHERE user = ?", login.User).Scan(&loginMysql.Password)

	if errr != nil{
		fmt.Print(errr)
	}

	logIn := utils.CompareHash(login.Password, loginMysql.Password)

	if logIn {
		c.JSON(200, gin.H{"pass": true})
	} else {
		c.JSON(500, errr)
	}



}

