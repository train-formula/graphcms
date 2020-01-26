package connections

import (
	"context"

	"github.com/train-formula/graphcms/models"
	"github.com/train-formula/graphcms/models/plan"
)

type PlanConnection struct {
	ResolveTotalCount TotalCountResolver
	Edges             []*plan.Plan `json:"edges"`
}

func (w *PlanConnection) TotalCount(ctx context.Context, obj *PlanConnection) (int, error) {

	return obj.ResolveTotalCount(ctx)
}

func (w *PlanConnection) PageInfo(ctx context.Context, obj *PlanConnection) (*models.PageInfo, error) {

	return GeneratePageInfo(len(obj.Edges), func() (string, string, error) {
		return obj.Edges[0].Cursor(), obj.Edges[len(obj.Edges)-1].Cursor(), nil
	})

}
