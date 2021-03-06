package plancall

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/train-formula/graphcms/database/plandb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/util"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewEditPlan(request generated.EditPlan, logger *zap.Logger, db pgxload.PgxLoader) *EditPlan {
	return &EditPlan{
		request: request,
		db:      db,
		logger:  logger.Named("EditPlan"),
	}
}

type EditPlan struct {
	request generated.EditPlan
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (g *EditPlan) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(g.request.Name, "If set, name must not be empty", true),
	}
}

func (g *EditPlan) Call(ctx context.Context) (*plan.Plan, error) {

	var finalPlan *plan.Plan

	err := pgxload.RunInTransaction(ctx, g.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		pln, err := plandb.GetPlanForUpdate(ctx, t, g.request.ID)
		if err != nil {
			if err == pgx.ErrNoRows {
				return gqlerror.Errorf("Plan does not exist")
			}

			g.logger.Error("Error retrieving plan", zap.Error(err),
				logging.UUID("planID", g.request.ID))
			return err
		}

		if g.request.Name != nil {
			pln.Name = strings.TrimSpace(*g.request.Name)
		}

		if g.request.Description != nil {
			pln.Description = util.TrimSpaceNotNil(g.request.Description.Value)
		}

		if g.request.RegistrationAvailable != nil {
			pln.RegistrationAvailable = *g.request.RegistrationAvailable
		}

		finalPlan, err = plandb.UpdatePlan(ctx, t, pln)
		if err != nil {
			g.logger.Error("Error updating plan", zap.Error(err),
				logging.UUID("planID", g.request.ID))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalPlan, nil
}
