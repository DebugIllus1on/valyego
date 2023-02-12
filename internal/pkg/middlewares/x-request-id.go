package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/valyego/internal/pkg/known"
)

func XRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.Request.Header.Get(known.XRequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx.Set(known.XRequestIDKey, requestID)
		ctx.Writer.Header().Set(known.XRequestIDKey, requestID)

		ctx.Next()
	}
}
