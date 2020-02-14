package organizationid

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/willtrking/pgxload"
)

const contextKey = "organizationID"

func GetContextLoader(ctx context.Context) *OrganizationIDLoader {
	return ctx.Value(contextKey).(*OrganizationIDLoader)
}

func AddContextLoader(ctx *gin.Context, db pgxload.PgxLoader) {

	dataloader.GinRegisterLoader(ctx, contextKey, NewLoader(ctx, db))

}
