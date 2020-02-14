package plandb

import (
	"context"

	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/willtrking/pgxload"
)

// Update plan, replace all fields with new row
func UpdatePlan(ctx context.Context, conn pgxload.PgxTxLoader, new plan.Plan) (*plan.Plan, error) {

	upd := pgxload.NewStructUpdate(database.TableName(new), new)

	upd = upd.WithReturning("*")

	updStmt, updParams, err := upd.GenerateExactUpdate(conn.Mapper(), "id")
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, updStmt, updParams...)
	if err != nil {
		return nil, err
	}

	var result plan.Plan

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

// Update plan schedule, replace all fields with new row
func UpdatePlanSchedule(ctx context.Context, conn pgxload.PgxTxLoader, new plan.PlanSchedule) (*plan.PlanSchedule, error) {

	upd := pgxload.NewStructUpdate(database.TableName(new), new)

	upd = upd.WithReturning("*")

	updStmt, updParams, err := upd.GenerateExactUpdate(conn.Mapper(), "id")
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, updStmt, updParams...)
	if err != nil {
		return nil, err
	}

	var result plan.PlanSchedule

	err = conn.Scanner(rows).Scan(&result)

	return &result, err

}
