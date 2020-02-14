package plancall

import (
	"context"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/plandb"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/util"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewCreatePlan(request generated.CreatePlan, logger *zap.Logger, db pgxload.PgxLoader) *CreatePlan {
	return &CreatePlan{
		request: request,
		db:      db,
		logger:  logger.Named("CreatePlan"),
	}
}

type CreatePlan struct {
	request generated.CreatePlan
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (g CreatePlan) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(g.request.Name, "Name must not be empty", true),
		validation.OrganizationExists(ctx, g.db, g.request.TrainerOrganizationID),
	}
}

func (g CreatePlan) Call(ctx context.Context) (*plan.Plan, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		g.logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var finalPlan *plan.Plan

	err = pgxload.RunInTransaction(ctx, g.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		err = validation.TagsAllExistForTrainer(ctx, t, g.request.TrainerOrganizationID, g.request.Tags)
		if err != nil {
			return err
		}

		newPlan := plan.Plan{
			ID: newUuid,

			TrainerOrganizationID: g.request.TrainerOrganizationID,

			Name:        strings.TrimSpace(g.request.Name),
			Description: util.TrimSpaceNotNil(g.request.Description),

			RegistrationAvailable: g.request.RegistrationAvailable,
			Archived:              false,
		}

		finalPlan, err = plandb.InsertPlan(ctx, t, newPlan)

		if err != nil {
			g.logger.Error("Failed to insert plan", zap.Error(err))
			return err
		}

		for _, tagUUID := range g.request.Tags {

			_, err := tagdb.TagPlan(ctx, t, tagUUID, g.request.TrainerOrganizationID, finalPlan.ID)
			if err != nil {
				g.logger.Error("Failed to tag plan", zap.Error(err))
				return err
			}
		}

		if g.request.Inventory != nil {

			err = plandb.SetPlanTotalInventory(ctx, t, finalPlan.ID, *g.request.Inventory)
			if err != nil {
				g.logger.Error("Failed to add inventory to plan", zap.Error(err))
				return err
			}
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalPlan, nil

}
