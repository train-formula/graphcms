package dataloader

import (
	"github.com/gin-gonic/gin"
)

func GinRegisterLoader(ctx *gin.Context, contextKey string, loader interface{}) {
	if _, exists := ctx.Get(contextKey); exists {
		panic("Cannot double-register data loader context key " + contextKey)
	}

	ctx.Set(contextKey, loader)
}
