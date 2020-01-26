package plandb

import (
	"context"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/plan"
)

func InsertPlan(ctx context.Context, conn database.Tx, new plan.Plan) (*plan.Plan, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

func InsertPlanSchedule(ctx context.Context, conn database.Tx, new plan.PlanSchedule) (*plan.PlanSchedule, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

// Set the total amount of inventory on a plan
func SetPlanTotalInventory(ctx context.Context, conn database.Tx, planID uuid.UUID, count int) error {

	invUuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	_, err = InsertOrUpdatePlanInventory(ctx, conn, plan.PlanInventory{
		ID:             invUuid,
		PlanID:         planID,
		PlanScheduleID: nil,
		TotalInventory: count,
	})

	if err != nil {
		return err
	}

	return nil

}

// Set the total amount of inventory on a plan schedule
func SetPlanScheduleTotalInventory(ctx context.Context, conn database.Tx, planID uuid.UUID, planScheduleID uuid.UUID, count int) error {

	invUuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	_, err = InsertOrUpdatePlanInventory(ctx, conn, plan.PlanInventory{
		ID:             invUuid,
		PlanID:         planID,
		PlanScheduleID: &planScheduleID,
		TotalInventory: count,
	})

	if err != nil {
		return err
	}

	return nil

}

// Will either insert a PlanInventory row with the specified plan + plan schedule ID, or update the values in the row
// Should be done in a transaction
func InsertOrUpdatePlanInventory(ctx context.Context, conn database.Tx, new plan.PlanInventory) (*plan.PlanInventory, error) {

	query, params, err := database.StructInsertStatement(new, "")

	if err != nil {
		return nil, err
	}

	query += " ON CONFLICT (\"plan_id\",COALESCE(\"plan_schedule_id\", '00000000-0000-0000-0000-000000000000'::uuid))"
	query += " DO UPDATE SET total_inventory = EXCLUDED.total_inventory"
	query += " RETURNING " + strings.Join(database.StructColumns(new, ""), ", ")

	var result plan.PlanInventory

	_, err = conn.QueryContext(ctx, query, &result, params...)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
