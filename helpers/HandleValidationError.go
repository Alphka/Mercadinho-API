package helpers

import (
	"errors"
	"fmt"
	"mercadinho/models"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(ctx *gin.Context, err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		var errorMessages []string

		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			param := fieldError.Param()
			kind := fieldError.Kind()

			if kind == reflect.Ptr {
				kind = fieldError.Type().Elem().Kind()
			}

			switch fieldError.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is required", fieldName))
			case "price":
				errorMessages = append(errorMessages, "Price must be greater than zero and have at most two decimal places")
			case "min":
				switch kind {
				case reflect.String:
					errorMessages = append(errorMessages, fmt.Sprintf("%s must have at least %s characters", fieldName, param))
				case reflect.Slice, reflect.Array, reflect.Map:
					errorMessages = append(errorMessages, fmt.Sprintf("%s must have at least %s items", fieldName, param))
				default:
					errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s", fieldName, param))
				}
			case "max":
				switch kind {
				case reflect.String:
					errorMessages = append(errorMessages, fmt.Sprintf("%s must have at most %s characters", fieldName, param))
				case reflect.Slice, reflect.Array, reflect.Map:
					errorMessages = append(errorMessages, fmt.Sprintf("%s must have at most %s items", fieldName, param))
				default:
					errorMessages = append(errorMessages, fmt.Sprintf("%s must be at most %s", fieldName, param))
				}
			case "len":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must have a length of %s characters", fieldName, param))
			case "gte":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be greater than or equal to %s", fieldName, param))
			case "lte":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be less than or equal to %s", fieldName, param))
			case "email":
				errorMessages = append(errorMessages, fmt.Sprintf("Invalid email format for %s", fieldName))
			case "uuid":
				errorMessages = append(errorMessages, fmt.Sprintf("Invalid UUID format for %s", fieldName))
			case "datetime", "date":
				errorMessages = append(errorMessages, fmt.Sprintf("Invalid datetime format for %s", fieldName))
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("Invalid value for %s", fieldName))
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
