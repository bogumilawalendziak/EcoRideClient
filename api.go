package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initAPI() {
	router := gin.Default()
	InitializeAPI(router)
	router.Run(":8081")
}

func InitializeAPI(router *gin.Engine) {
	router.POST("/reserve", handleReserveRequest)
}

func handleReserveRequest(c *gin.Context) {
	var request ReserveRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sendReserveRequest(request)

	c.JSON(http.StatusOK, nil)
}
