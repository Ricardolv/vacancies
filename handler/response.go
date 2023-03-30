package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccessOK(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

func sendSuccessCreated(ctx *gin.Context, id int) {
	ctx.Header("Content-Type", "application/json")
	ctx.Header("location", fmt.Sprintf("localhost:8080/api/v1/openings/%v", id))
	ctx.AbortWithStatus(http.StatusCreated)
}
