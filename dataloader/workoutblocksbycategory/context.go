package workoutblocksbycategory

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/dataloader"
)

const contextKey = "workoutBlocksByCategory"

func GetContextLoader(ctx context.Context) *WorkoutBlocksByCategoryLoader {
	return ctx.Value(contextKey).(*WorkoutBlocksByCategoryLoader)
}

func AddContextLoader(ctx *gin.Context, db *pg.DB) {
	dataloader.GinRegisterLoader(ctx, db, contextKey, NewLoader(ctx, db))
}
