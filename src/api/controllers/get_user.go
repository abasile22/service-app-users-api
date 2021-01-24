package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

type MySql struct{
	User string
	Password string
}

func GetUser(c *gin.Context) {

	id := c.Param("id")

	var login MySql

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/services_app_users")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// See "Important settings" section.
	// perform a db.Query insert
	errr := db.QueryRow("SELECT password FROM users WHERE user = ?", id).Scan(&login.Password)

	if errr != nil{
		fmt.Print(errr)
	}

	fmt.Print(login.Password)

	c.JSON(200, gin.H{"pass": login.Password})

}

