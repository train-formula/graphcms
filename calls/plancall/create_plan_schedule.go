package plancall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/plandb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/util"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func NewCreatePlanSchedule(request generated.CreatePlanSchedule, logger *zap.Logger, db *pg.DB) *CreatePlanSchedule {
	return &CreatePlanSchedule{
		request: request,
		db:      db,
		logger:  logger.Named("CreatePlanSchedule"),
	}
}

type CreatePlanSchedule struct {
	request generated.CreatePlanSchedule
	db      *pg.DB
	logger  *zap.Logger
}

func (g CreatePlanSchedule) Validate(ctx context.Context) []validation.ValidatorFunc {

	funcs := []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(g.request.Name, "If specified, name must not be empty", true),
		validation.CheckIntGTE(g.request.PricePerInterval, 0, "Price per interval must be >= 0"),
		validation.CheckIntIsNilOrGTE(g.request.PriceMarkedDownFrom, 0, "If set, price marked down from must be >= 0"),
		validation.OrganizationExists(ctx, g.db, g.request.TrainerOrganizationID),
		validation.PlanExists(ctx, g.db, g.request.PlanID),
	}

	if g.request.PaymentInterval != nil {
		fieldName := "paymentInterval"
		funcs = append(funcs, validation.CheckDiurnalIntervalInput(*g.request.PaymentInterval, 1, &fieldName))
	} else {
		funcs = append(funcs, validation.ImmediateErrorValidator("Payment interval is required"))
	}

	if g.request.DurationInterval != nil {
		fieldName := "durationInterval"
		funcs = append(funcs, validation.CheckDiurnalIntervalInput(*g.request.DurationInterval, 1, &fieldName))
	}

	return funcs
}

func (g CreatePlanSchedule) Call(ctx context.Context) (*plan.PlanSchedule, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		g.logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var finalPlanSchedule *plan.PlanSchedule

	err = g.db.RunInTransaction(func(t *pg.Tx) error {

		newSchedule := plan.PlanSchedule{
			ID: newUuid,

			TrainerOrganizationID: g.request.TrainerOrganizationID,
			PlanID:                g.request.PlanID,

			Name:        util.TrimSpaceNotNil(g.request.Name),
			Description: util.TrimSpaceNotNil(g.request.Description),

			PaymentInterval:       *g.request.PaymentInterval.Interval,
			PaymentIntervalCount:  g.request.PaymentInterval.Count,
			PricePerInterval:      g.request.PricePerInterval,
			PriceMarkedDownFrom:   g.request.PriceMarkedDownFrom,
			DurationInterval:      g.request.DurationInterval.Interval,
			DurationIntervalCount: &g.request.DurationInterval.Count,

			RegistrationAvailable: g.request.RegistrationAvailable,
			Archived:              false,
		}

		finalPlanSchedule, err = plandb.InsertPlanSchedule(ctx, t, newSchedule)

		if err != nil {
			g.logger.Error("Failed to insert plan schedule", zap.Error(err))
			return err
		}

		if g.request.Inventory != nil {

			err = plandb.SetPlanScheduleTotalInventory(ctx, t, finalPlanSchedule.PlanID, finalPlanSchedule.ID, *g.request.Inventory)
			if err != nil {
				g.logger.Error("Failed to add inventory to plan schedule", zap.Error(err))
				return err
			}
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalPlanSchedule, nil

}
