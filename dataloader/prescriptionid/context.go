package prescriptionid

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/willtrking/pgxload"
)

const contextKey = "prescriptionID"

func GetContextLoader(ctx context.Context) *PrescriptionIDLoader {
	return ctx.Value(contextKey).(*PrescriptionIDLoader)
}

func AddContextLoader(ctx *gin.Context, db pgxload.PgxLoader) {
	dataloader.GinRegisterLoader(ctx, contextKey, NewLoader(ctx, db))
}
