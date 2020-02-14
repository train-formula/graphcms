//go:generate go run github.com/vektah/dataloaden TagIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/tag.tag

package tagid

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/willtrking/pgxload"
)

func NewLoader(ctx context.Context, db pgxload.PgxLoader) *TagIDLoader {

	return NewTagIDLoader(TagIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*tag.Tag, []error) {

			tags, err := tagdb.GetTags(ctx, db, keys)

			result := make([]*tag.Tag, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			tagsMap := make(map[uuid.UUID]*tag.Tag)

			for _, o := range tags {
				tagsMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasTag := tagsMap[k]; hasTag {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
