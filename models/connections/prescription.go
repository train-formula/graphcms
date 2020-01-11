package connections

import (
	"context"

	"github.com/train-formula/graphcms/models"
	"github.com/train-formula/graphcms/models/workout"
)

type PrescriptionConnection struct {
	ResolveTotalCount TotalCountResolver
	Edges             []*workout.Prescription `json:"edges"`
}

func (w *PrescriptionConnection) TotalCount(ctx context.Context, obj *PrescriptionConnection) (int, error) {

	return obj.ResolveTotalCount(ctx)
}

func (w *PrescriptionConnection) PageInfo(ctx context.Context, obj *PrescriptionConnection) (*models.PageInfo, error) {

	return GeneratePageInfo(len(obj.Edges), func() (string, string, error) {
		return obj.Edges[0].Cursor(), obj.Edges[len(obj.Edges)-1].Cursor(), nil
	})

}
