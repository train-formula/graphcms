package workoutblockid

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/dataloader"
)

const contextKey = "workoutBlockID"

func GetContextLoader(ctx context.Context) *WorkoutBlockIDLoader {
	return ctx.Value(contextKey).(*WorkoutBlockIDLoader)
}

func AddContextLoader(ctx *gin.Context, db *pg.DB) {
	dataloader.GinRegisterLoader(ctx, db, contextKey, NewLoader(ctx, db))
}
