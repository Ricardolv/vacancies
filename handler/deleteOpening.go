package handler

import (
	"fmt"
	"net/http"

	"github.com/Ricardolv/vacancies/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 204 {object} ErrorResponse
// @Router /openings [delete]
func DeleteOpeningHandler(ctx *gin.Context) {

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

	// Delete the opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting opening whit id: %s", id))
		return
	}

	sendSuccessOK(ctx, "deleted-opening", opening)
}
