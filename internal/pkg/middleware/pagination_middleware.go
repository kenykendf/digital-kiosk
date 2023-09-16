package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationMiddleware(defaultOffset int, defaultLimit int, defaultAsc int) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		offset, err := strconv.Atoi(ctx.Query("offset"))
		if err != nil {
			offset = defaultOffset
		}

		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			limit = defaultLimit
		}

		asc, err := strconv.Atoi(ctx.Query("asc"))
		if err != nil {
			asc = defaultAsc
		}

		ctx.Set("offset", offset)
		ctx.Set("limit", limit)
		ctx.Set("asc", asc)

		ctx.Next()
	}

}
