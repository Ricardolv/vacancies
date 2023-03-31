package handler

import (
	"net/http"

	"github.com/Ricardolv/vacancies/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 204
// @Router /openings [get]
func ShowOpeningHandler(ctx *gin.Context) {

	// Get the user in ctx
	id := ctx.Query("id")
	if len(id) == 0 {
		sendError(ctx, http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find the opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNoContent, "")
		return
	}

	sendSuccessOK(ctx, "show-opening", opening)
}
