package plancall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/plandb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/util"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewEditPlanSchedule(request generated.EditPlanSchedule, logger *zap.Logger, db *pg.DB) *EditPlanSchedule {
	return &EditPlanSchedule{
		request: request,
		db:      db,
		logger:  logger.Named("EditPlanSchedule"),
	}
}

type EditPlanSchedule struct {
	request generated.EditPlanSchedule
	db      *pg.DB
	logger  *zap.Logger
}

func (g EditPlanSchedule) Validate(ctx context.Context) []validation.ValidatorFunc {

	var funcs []validation.ValidatorFunc

	if g.request.Name != nil {
		funcs = append(funcs, validation.CheckStringNilOrIsNotEmpty(g.request.Name.Value, "If set, name must not be empty", true))
	}

	if g.request.PriceMarkedDownFrom != nil {
		funcs = append(funcs, validation.CheckIntIsNilOrGTE(g.request.PriceMarkedDownFrom.Value, 0, "If set, price marked down from must be >= 0"))
	}

	return funcs
}

func (g EditPlanSchedule) Call(ctx context.Context) (*plan.PlanSchedule, error) {

	var finalPlanSchedule *plan.PlanSchedule

	err := g.db.RunInTransaction(func(t *pg.Tx) error {

		planSchedule, err := plandb.GetPlanScheduleForUpdate(ctx, t, g.request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Plan schedule does not exist")
			}

			g.logger.Error("Error retrieving plan schedule", zap.Error(err),
				logging.UUID("planScheduleID", g.request.ID))
			return err
		}

		if g.request.Name != nil {
			planSchedule.Name = util.TrimSpaceNotNil(g.request.Name.Value)
		}

		if g.request.Description != nil {
			planSchedule.Description = util.TrimSpaceNotNil(g.request.Description.Value)
		}

		if g.request.PriceMarkedDownFrom != nil {
			planSchedule.PriceMarkedDownFrom = g.request.PriceMarkedDownFrom.Value
		}

		if g.request.RegistrationAvailable != nil {
			planSchedule.RegistrationAvailable = *g.request.RegistrationAvailable
		}

		finalPlanSchedule, err = plandb.UpdatePlanSchedule(ctx, t, planSchedule)
		if err != nil {
			g.logger.Error("Error updating plan schedule", zap.Error(err),
				logging.UUID("planScheduleID", g.request.ID))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalPlanSchedule, nil
}
