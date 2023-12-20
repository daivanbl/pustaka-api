package main

import (
	"pustaka-api/config"
	"pustaka-api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()

	router := gin.Default()

	route.ApiRoute(router, db)

	router.Run(":8888")
}
