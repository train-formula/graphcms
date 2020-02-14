package blockexercisesbyblock

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/willtrking/pgxload"
)

const contextKey = "blockExercisesByBlock"

func GetContextLoader(ctx context.Context) *BlockExercisesByBlockLoader {
	return ctx.Value(contextKey).(*BlockExercisesByBlockLoader)
}

func AddContextLoader(ctx *gin.Context, db pgxload.PgxLoader) {
	dataloader.GinRegisterLoader(ctx, contextKey, NewLoader(ctx, db))
}
