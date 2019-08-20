package connections

import (
	"context"

	"github.com/onsi/gomega/format"
	"github.com/train-formula/graphcms/models"
	"github.com/train-formula/graphcms/models/workout"
)

type WorkoutProgramConnection struct {
	ResolveTotalCount func(ctx format.Ctx) (int, error)
	Edges             []*workout.WorkoutProgram `json:"edges"`
}

func (w *WorkoutProgramConnection) TotalCount(ctx context.Context, obj *WorkoutProgramConnection) (int, error) {

	return obj.ResolveTotalCount(ctx)
}

func (w *WorkoutProgramConnection) PageInfo(ctx context.Context, obj *WorkoutProgramConnection) (*models.PageInfo, error) {

	if len(obj.Edges) > 0 {
		return &models.PageInfo{
			StartCursor: obj.Edges[0].Cursor(),
			EndCursor:   obj.Edges[len(obj.Edges)-1].Cursor(),
		}, nil
	}

	return &models.PageInfo{}, nil

}
