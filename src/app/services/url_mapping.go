package services

import (
	"github.com/abasile22/service-app-api/src/app/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Start() {
	router := gin.Default()
	router.GET("/ping", Ping)

	router.POST("/user/new-user", users.CreateNewUser)
	router.POST("/user/update-user/:id", users.UpdateUser)
	router.GET("/user/get-user/:id", users.GetUser)
	router.POST("/user/log-in", users.LogInUser)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.

	// router.Run(":3000") for a hard coded port

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		router.Run(":8080")
	}
	log.Fatal(http.ListenAndServe(port, router))
}