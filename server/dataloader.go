package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/dataloader/organizationid"
	"github.com/train-formula/graphcms/dataloader/tagbytag"
	"github.com/train-formula/graphcms/dataloader/tagid"
)

func RegisterLoaders(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		organizationid.AddContextLoader(c, db)
		tagid.AddContextLoader(c, db)
		tagbytag.AddContextLoader(c, db)

		c.Request = c.Request.WithContext(c)

		c.Next()

	}
}
