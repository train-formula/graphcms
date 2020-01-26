package plandb

import (
	"context"

	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/plan"
)

// Update plan, replace all fields with new row
func UpdatePlan(ctx context.Context, conn database.Tx, new plan.Plan) (*plan.Plan, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Update()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

// Update plan schedule, replace all fields with new row
func UpdatePlanSchedule(ctx context.Context, conn database.Tx, new plan.PlanSchedule) (*plan.PlanSchedule, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Update()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}
