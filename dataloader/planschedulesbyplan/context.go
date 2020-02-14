package planschedulesbyplan

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/willtrking/pgxload"
)

const contextKey = "planSchedulesByPlan"

func GetContextLoader(ctx context.Context) *PlanSchedulesByPlanLoader {
	return ctx.Value(contextKey).(*PlanSchedulesByPlanLoader)
}

func AddContextLoader(ctx *gin.Context, db pgxload.PgxLoader) {
	dataloader.GinRegisterLoader(ctx, contextKey, NewLoader(ctx, db))
}
