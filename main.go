package main

import (
	"marvel/adapter"
	"marvel/service"
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

	router.GET("/characters/:id", adapter.GetCharacter(service.GetTs, http.DefaultClient))

	router.StaticFile("/characters", "./static/allcharacters.json")

	router.Run("0.0.0.0:8080")
}
