package handler

import (
	"net/http"

	"github.com/Ricardolv/vacancies/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {

	// Get the user in ctx
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

	sendSuccessOK(ctx, "show-opening", opening)
}
