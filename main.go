package main

import (
	"marvel/adapter"
	"marvel/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title Marvel Character API
// @version 1.0
// @description Ultimate Marvel Character API

// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()
	router.GET("/characters/:id", adapter.GetCharacter(controller.GetTs, http.DefaultClient))

	router.Run("0.0.0.0:8080")
}
