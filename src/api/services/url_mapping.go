package services

import (
	"github.com/abasile22/service-app-api/src/api/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Start() {
	router := gin.Default()
	router.GET("/ping", Ping)

	router.POST("/users/new-user", users.CreateNewUser)
	router.POST("/users/update-user/:id", users.UpdateUser)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.

	// router.Run(":3000") for a hard coded port

	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, router))
}