package blockexercisesbyblock

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/dataloader"
)

const contextKey = "blockExercisesByBlock"

func GetContextLoader(ctx context.Context) *BlockExercisesByBlockLoader {
	return ctx.Value(contextKey).(*BlockExercisesByBlockLoader)
}

func AddContextLoader(ctx *gin.Context, db *pg.DB) {
	dataloader.GinRegisterLoader(ctx, db, contextKey, NewLoader(ctx, db))
}
