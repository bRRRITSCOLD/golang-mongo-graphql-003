package dataloaders

import (
	"context"

	"github.com/gin-gonic/gin"
)

// func GinContextToContextMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx := context.WithValue(c.Request.Context(), GIN_CONTEXT_KEY, c)
// 		c.Request = c.Request.WithContext(ctx)
// 		c.Next()
// 	}
// }

// Middleware stores Loaders as a request-scoped context value.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		loaders := newLoaders(ctx)
		augmentedCtx := context.WithValue(ctx, key, loaders)
		c.Request = c.Request.WithContext(augmentedCtx)
		c.Next()
	}
}
