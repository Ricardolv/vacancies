package handler

import (
	"net/http"

	"github.com/Ricardolv/vacancies/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}
	// Find the opening
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "")
		return
	}

	sendSuccessOK(ctx, "list-opening", openings)
}
