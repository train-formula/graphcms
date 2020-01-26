package planschedulesbyplan

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/dataloader"
)

const contextKey = "planSchedulesByPlan"

func GetContextLoader(ctx context.Context) *PlanSchedulesByPlanLoader {
	return ctx.Value(contextKey).(*PlanSchedulesByPlanLoader)
}

func AddContextLoader(ctx *gin.Context, db *pg.DB) {
	dataloader.GinRegisterLoader(ctx, db, contextKey, NewLoader(ctx, db))
}
