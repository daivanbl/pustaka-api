package route

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PokemonRoute(pokes *gin.RouterGroup, db *gorm.DB) {
	pokes.GET("/:regionName", handler.GetListPokemonByRegion)
}
