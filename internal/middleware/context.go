package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ContextKey string

const (
	GIN_CONTEXT_KEY ContextKey = "GinContextKey"
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GIN_CONTEXT_KEY, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
