//go:generate go run github.com/vektah/dataloaden TagByTagLoader github.com/train-formula/graphcms/database/tagdb.TagByTag *github.com/train-formula/graphcms/models/tag.tag

package tagbytag

import (
	"context"

	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/willtrking/pgxload"
)

func NewLoader(ctx context.Context, db pgxload.PgxLoader) *TagByTagLoader {

	return NewTagByTagLoader(TagByTagLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []tagdb.TagByTag) ([]*tag.Tag, []error) {

			tags, err := tagdb.GetTagsByTag(ctx, db, keys)

			result := make([]*tag.Tag, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			tagsMap := make(map[tagdb.TagByTag]*tag.Tag)

			for _, o := range tags {
				tagsMap[tagdb.TagByTag{
					Tag:                   o.Tag,
					TrainerOrganizationID: o.TrainerOrganizationID,
				}.Stable()] = o
			}

			for i, k := range keys {
				if tg, hasTag := tagsMap[k.Stable()]; hasTag {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
