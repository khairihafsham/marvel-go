package adapter

import (
	"marvel/controller"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET Character godoc
// @Summary Get marvel character by id
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} controller.Character
// @Failure 404 {string} string Character not found
// @Failure 500 {string} string Error while processing request
// @Router /characters/{id} [get]
func GetCharacter(getTs func() int64, client *http.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		result := controller.GetMarvelCharacter(getTs, id, client)

		if result.StatusCode != 200 {
			c.AbortWithStatus(result.StatusCode)
			return
		}

		c.IndentedJSON(result.StatusCode, result.Character)
	}
}
