package handler

import (
	"fmt"
	"net/http"

	"github.com/Ricardolv/vacancies/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {

	id := ctx.Query("id")
	if len(id) == 0 {
		sendError(ctx, http.StatusBadRequest,
			errParamIsRequest("id", "queryParameter").Error())
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
