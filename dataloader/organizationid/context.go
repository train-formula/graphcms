package organizationid

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/dataloader"
)

const contextKey = "organizationID"

func GetContextLoader(ctx context.Context) *OrganizationIDLoader {
	return ctx.Value(contextKey).(*OrganizationIDLoader)
}

func AddContextLoader(ctx *gin.Context, db *pg.DB) {

	dataloader.GinRegisterLoader(ctx, db, contextKey, NewLoader(ctx, db))

}
