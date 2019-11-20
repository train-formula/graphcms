package connections

import (
	"context"

	"github.com/train-formula/graphcms/models"
	"github.com/train-formula/graphcms/models/workout"
)

type WorkoutCategoryConnection struct {
	ResolveTotalCount TotalCountResolver
	Edges             []*workout.WorkoutCategory `json:"edges"`
}

func (w *WorkoutCategoryConnection) TotalCount(ctx context.Context, obj *WorkoutCategoryConnection) (int, error) {

	return obj.ResolveTotalCount(ctx)
}

func (w *WorkoutCategoryConnection) PageInfo(ctx context.Context, obj *WorkoutCategoryConnection) (*models.PageInfo, error) {

	return GeneratePageInfo(len(obj.Edges), func() (string, string, error) {
		return obj.Edges[0].Cursor(), obj.Edges[len(obj.Edges)-1].Cursor(), nil
	})

}
