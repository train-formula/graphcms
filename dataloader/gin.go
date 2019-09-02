package dataloader

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

func GinRegisterLoader(ctx *gin.Context, db *pg.DB, contextKey string, loader interface{}) {
	if _, exists := ctx.Get(contextKey); exists {
		panic("Cannot double-register data loader context key " + contextKey)
	}

	ctx.Set(contextKey, loader)
}
