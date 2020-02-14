package workoutblocksbycategory

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/willtrking/pgxload"
)

const contextKey = "workoutBlocksByCategory"

func GetContextLoader(ctx context.Context) *WorkoutBlocksByCategoryLoader {
	return ctx.Value(contextKey).(*WorkoutBlocksByCategoryLoader)
}

func AddContextLoader(ctx *gin.Context, db pgxload.PgxLoader) {
	dataloader.GinRegisterLoader(ctx, contextKey, NewLoader(ctx, db))
}
