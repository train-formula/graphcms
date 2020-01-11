//go:generate go run github.com/vektah/dataloaden TagsByObjectLoader github.com/train-formula/graphcms/database/tagdb.TagsByObject []*github.com/train-formula/graphcms/models/tag.tag

package tagsbyobject

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/tag"
)

func NewLoader(ctx context.Context, db *pg.DB) *TagsByObjectLoader {

	return NewTagsByObjectLoader(TagsByObjectLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []tagdb.TagsByObject) ([][]*tag.Tag, []error) {

			tags, err := tagdb.GetTagsByObject(ctx, db, keys)

			result := make([][]*tag.Tag, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			for idx, k := range keys {

				if tg, hasTg := tags[k]; hasTg {
					result[idx] = tg
				} else {
					result[idx] = make([]*tag.Tag, 0, 0)
				}
			}

			return result, errors
		},
	})
}
