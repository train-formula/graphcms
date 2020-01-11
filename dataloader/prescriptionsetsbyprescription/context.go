package prescriptionsetsbyprescription

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/dataloader"
)

const contextKey = "prescriptionSetsByPrescription"

func GetContextLoader(ctx context.Context) *PrescriptionSetsByPrescriptionLoader {
	return ctx.Value(contextKey).(*PrescriptionSetsByPrescriptionLoader)
}

func AddContextLoader(ctx *gin.Context, db *pg.DB) {
	dataloader.GinRegisterLoader(ctx, db, contextKey, NewLoader(ctx, db))
}
