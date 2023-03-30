package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	router := gin.Default()

	// Initialize routes
	InitializeRoutes(router)

	router.Run(":8080")
}
