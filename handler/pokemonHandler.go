package handler

import (
	"net/http"
	"pustaka-api/service"

	"github.com/gin-gonic/gin"
)

func GetListPokemonByRegion(c *gin.Context) {
	regionName := c.Param("regionName")

	response := service.FindPokemonsByRegion(regionName)

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   response,
	})
}
