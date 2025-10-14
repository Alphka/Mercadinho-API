package helpers

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidatePagination(ctx *gin.Context, totalItems int) (page, pageSize, totalPages int) {
	page, _ = strconv.Atoi(ctx.Query("page"))
	pageSize, _ = strconv.Atoi(ctx.Query("pageSize"))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 20
	}

	totalPages = int(math.Ceil(float64(totalItems) / float64(pageSize)))

	return
}
