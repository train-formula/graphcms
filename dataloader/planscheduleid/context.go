package planscheduleid

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/willtrking/pgxload"
)

const contextKey = "planScheduleID"

func GetContextLoader(ctx context.Context) *PlanScheduleIDLoader {
	return ctx.Value(contextKey).(*PlanScheduleIDLoader)
}

func AddContextLoader(ctx *gin.Context, db pgxload.PgxLoader) {
	dataloader.GinRegisterLoader(ctx, contextKey, NewLoader(ctx, db))
}
