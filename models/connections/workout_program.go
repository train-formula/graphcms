package connections

import (
	"context"

	"github.com/train-formula/graphcms/models"
	"github.com/train-formula/graphcms/models/workout"
)

type WorkoutProgramConnection struct {
	ResolveTotalCount TotalCountResolver
	Edges             []*workout.WorkoutProgram `json:"edges"`
}

func (w *WorkoutProgramConnection) TotalCount(ctx context.Context, obj *WorkoutProgramConnection) (int, error) {

	return obj.ResolveTotalCount(ctx)
}

func (w *WorkoutProgramConnection) PageInfo(ctx context.Context, obj *WorkoutProgramConnection) (*models.PageInfo, error) {

	return GeneratePageInfo(len(obj.Edges), func() (string, string, error) {
		return obj.Edges[0].Cursor(), obj.Edges[len(obj.Edges)-1].Cursor(), nil
	})

}
