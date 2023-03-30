package router

import (
	"github.com/Ricardolv/vacancies/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.ShowOpeningHandler)

		v1.POST("/openings", handler.CreateOpeningHandler)

		v1.DELETE("/openings", handler.DeleteOpeningHandler)

		v1.PUT("/openings", handler.UpdateOpeningHandler)

		v1.GET("/openings/list", handler.ListOpeningsHandler)
	}
}
