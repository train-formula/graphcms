package connections

import (
	"context"

	"github.com/train-formula/graphcms/models"
	"github.com/train-formula/graphcms/models/tag"
)

type TagConnection struct {
	ResolveTotalCount TotalCountResolver
	Edges             []*tag.Tag `json:"edges"`
}

func (w *TagConnection) TotalCount(ctx context.Context, obj *TagConnection) (int, error) {

	return obj.ResolveTotalCount(ctx)
}

func (w *TagConnection) PageInfo(ctx context.Context, obj *TagConnection) (*models.PageInfo, error) {

	return GeneratePageInfo(len(obj.Edges), func() (string, string, error) {
		return obj.Edges[0].Cursor(), obj.Edges[len(obj.Edges)-1].Cursor(), nil
	})

}
