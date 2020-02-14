package prescriptionsetsbyprescription

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/willtrking/pgxload"
)

const contextKey = "prescriptionSetsByPrescription"

func GetContextLoader(ctx context.Context) *PrescriptionSetsByPrescriptionLoader {
	return ctx.Value(contextKey).(*PrescriptionSetsByPrescriptionLoader)
}

func AddContextLoader(ctx *gin.Context, db pgxload.PgxLoader) {
	dataloader.GinRegisterLoader(ctx, contextKey, NewLoader(ctx, db))
}
