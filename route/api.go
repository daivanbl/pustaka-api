package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiRoute(router *gin.Engine, db *gorm.DB) {
	BookRoute(router.Group("/books"), db)
}
