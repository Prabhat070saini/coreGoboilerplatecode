package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

// Default error message (same as NestJS)
const defaultMessage = "ERROR"

// SendRestResponse mirrors the NestJS sendRestResponse function
func SendRestResponse[T any](ctx *gin.Context, output *ServiceOutput[T]) {
	if output == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": defaultMessage,
		})
		return
	}

	// Handle Exception case
	if output.Exception != nil {
		ex := output.Exception
		statusCode := ex.HttpStatusCode
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}

		ctx.JSON(statusCode, gin.H{
			"code":    ex.Code,
			"message": fallbackIfEmpty(ex.Message, defaultMessage),
			"data":nil,
		})
		return
	}

	// Handle Success case
	if output.Success != nil {
		success := output.Success
		statusCode := success.HttpStatusCode
		if statusCode == 0 {
			statusCode = http.StatusOK
		}

		ctx.JSON(statusCode, gin.H{
			"code":    success.Code,
			"message": fallbackIfEmpty(success.Message, "Success"),
			"data":    success.Data,
		})
		return
	}

	// Default fallback response
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    nil,
	})
}

func fallbackIfEmpty(preferred, fallback string) string {
	if preferred != "" {
		return preferred
	}
	return fallback
}
