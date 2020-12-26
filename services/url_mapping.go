package services

import "github.com/gin-gonic/gin"

func App() {
	router := gin.Default()
	router.GET("/ping", Ping)
	router.POST("/new-user", NewUser)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.

	// router.Run(":3000") for a hard coded port

	router.Run(":8080")
}
