package router

import (
	docs "github.com/Ricardolv/vacancies/docs"
	"github.com/Ricardolv/vacancies/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/openings", handler.ShowOpeningHandler)

		v1.POST("/openings", handler.CreateOpeningHandler)

		v1.DELETE("/openings", handler.DeleteOpeningHandler)

		v1.PUT("/openings", handler.UpdateOpeningHandler)

		v1.GET("/openings/list", handler.ListOpeningsHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
