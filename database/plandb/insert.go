package plandb

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/willtrking/pgxload"
)

func InsertPlan(ctx context.Context, conn pgxload.PgxTxLoader, new plan.Plan) (*plan.Plan, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result plan.Plan

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

func InsertPlanSchedule(ctx context.Context, conn pgxload.PgxTxLoader, new plan.PlanSchedule) (*plan.PlanSchedule, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result plan.PlanSchedule

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

// Set the total amount of inventory on a plan
func SetPlanTotalInventory(ctx context.Context, conn pgxload.PgxTxLoader, planID uuid.UUID, count int) error {

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
func SetPlanScheduleTotalInventory(ctx context.Context, conn pgxload.PgxTxLoader, planID uuid.UUID, planScheduleID uuid.UUID, count int) error {

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
func InsertOrUpdatePlanInventory(ctx context.Context, conn pgxload.PgxTxLoader, new plan.PlanInventory) (*plan.PlanInventory, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*").WithConflict("(\"plan_id\",COALESCE(\"plan_schedule_id\", '00000000-0000-0000-0000-000000000000'::uuid)) DO UPDATE SET total_inventory = EXCLUDED.total_inventory")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result plan.PlanInventory

	err = conn.Scanner(rows).Scan(&result)

	return &result, err

}
