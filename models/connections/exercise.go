package connections

import (
	"context"

	"github.com/train-formula/graphcms/models"
	"github.com/train-formula/graphcms/models/workout"
)

type ExerciseConnection struct {
	ResolveTotalCount TotalCountResolver
	Edges             []*workout.Exercise `json:"edges"`
}

func (w *ExerciseConnection) TotalCount(ctx context.Context, obj *ExerciseConnection) (int, error) {

	return obj.ResolveTotalCount(ctx)
}

func (w *ExerciseConnection) PageInfo(ctx context.Context, obj *ExerciseConnection) (*models.PageInfo, error) {

	return GeneratePageInfo(len(obj.Edges), func() (string, string, error) {
		return obj.Edges[0].Cursor(), obj.Edges[len(obj.Edges)-1].Cursor(), nil
	})

}
