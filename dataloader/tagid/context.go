package tagid

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/willtrking/pgxload"
)

const contextKey = "tagID"

func GetContextLoader(ctx context.Context) *TagIDLoader {
	return ctx.Value(contextKey).(*TagIDLoader)
}

func AddContextLoader(ctx *gin.Context, db pgxload.PgxLoader) {
	dataloader.GinRegisterLoader(ctx, contextKey, NewLoader(ctx, db))
}
