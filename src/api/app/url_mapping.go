package app

import (
	"github.com/abasile22/service-app-api/src/api/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Start() {
	router := gin.Default()
	router.GET("/ping", controllers.Ping)

	router.POST("/user/new-user", controllers.CreateNewUser)
	router.POST("/user/update-user/:id", controllers.UpdateUser)
	router.GET("/user/get-user/:id", controllers.GetUser)
	router.POST("/user/log-in", controllers.LogInUser)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.

	// router.Run(":3000") for a hard coded port

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		router.Run(":8080")
	}
	log.Fatal(http.ListenAndServe(port, router))
}