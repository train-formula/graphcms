package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/dataloader/blockexercisesbyblock"
	"github.com/train-formula/graphcms/dataloader/exerciseid"
	"github.com/train-formula/graphcms/dataloader/organizationid"
	"github.com/train-formula/graphcms/dataloader/prescriptionid"
	"github.com/train-formula/graphcms/dataloader/tagbytag"
	"github.com/train-formula/graphcms/dataloader/tagid"
	"github.com/train-formula/graphcms/dataloader/tagsbyobject"
	"github.com/train-formula/graphcms/dataloader/unitid"
	"github.com/train-formula/graphcms/dataloader/workoutblockid"
	"github.com/train-formula/graphcms/dataloader/workoutblocksbycategory"
	"github.com/train-formula/graphcms/dataloader/workoutcategoriesbyworkout"
	"github.com/train-formula/graphcms/dataloader/workoutcategoryid"
	"github.com/train-formula/graphcms/dataloader/workoutid"
)

func RegisterLoaders(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		organizationid.AddContextLoader(c, db)
		tagid.AddContextLoader(c, db)
		prescriptionid.AddContextLoader(c, db)
		tagbytag.AddContextLoader(c, db)
		tagsbyobject.AddContextLoader(c, db)
		workoutcategoryid.AddContextLoader(c, db)
		unitid.AddContextLoader(c, db)
		workoutblocksbycategory.AddContextLoader(c, db)
		workoutid.AddContextLoader(c, db)
		workoutblockid.AddContextLoader(c, db)
		workoutcategoriesbyworkout.AddContextLoader(c, db)
		exerciseid.AddContextLoader(c, db)
		blockexercisesbyblock.AddContextLoader(c, db)

		c.Request = c.Request.WithContext(c)

		c.Next()

	}
}
