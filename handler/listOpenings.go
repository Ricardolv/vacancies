package handler

import (
	"net/http"

	"github.com/Ricardolv/vacancies/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/list [get]
func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}
	// Find the opening
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "")
		return
	}

	sendSuccessOK(ctx, "list-opening", openings)
}
