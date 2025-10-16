package helpers

import (
	"errors"
	"fmt"
	"mercadinho/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(ctx *gin.Context, err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		var errorMessages []string

		for _, fieldErr := range validationErrors {
			switch fieldErr.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is required", fieldErr.Field()))
			case "price":
				errorMessages = append(errorMessages, "Price must be greater than zero and have at most two decimal places")
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must have at least %s characters", fieldErr.Field(), fieldErr.Param()))
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("Invalid value for %s", fieldErr.Field()))
			}
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Error: strings.Join(errorMessages, ", "),
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
		Error: "Invalid JSON payload",
	})
}
