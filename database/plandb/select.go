package plandb

import (
	"context"
	"strings"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/plan"
)

// Retrieves individual plan's by their IDs
func GetPlans(ctx context.Context, conn database.Conn, ids []uuid.UUID) ([]*plan.Plan, error) {

	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*plan.Plan

	query := "SELECT * FROM " + database.TableName(plan.Plan{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}

// Retrieves individual plan schedules by their IDs
func GetPlanSchedules(ctx context.Context, conn database.Conn, ids []uuid.UUID) ([]*plan.PlanSchedule, error) {

	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*plan.PlanSchedule

	query := "SELECT * FROM " + database.TableName(plan.PlanSchedule{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}

// Retrieves plan schedules for plans mapped by their plan IDs
func GetSchedulesForPlans(ctx context.Context, conn database.Conn, planIDs []uuid.UUID) (map[uuid.UUID][]*plan.PlanSchedule, error) {

	results := make(map[uuid.UUID][]*plan.PlanSchedule)

	if len(planIDs) <= 0 {
		return results, nil
	}

	query := "SELECT * FROM " + database.TableName(plan.PlanSchedule{}) + " WHERE "

	var params []interface{}

	var queryResults []*plan.PlanSchedule

	for _, id := range planIDs {
		query += "plan_id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	_, err := conn.QueryContext(ctx, &queryResults, query, params...)

	if err != nil {
		return nil, err
	}

	for _, queryResult := range queryResults {
		results[queryResult.PlanID] = append(results[queryResult.PlanID], queryResult)
	}

	return results, err
}

// Retrieves a plan by its id
func GetPlan(ctx context.Context, conn database.Conn, id uuid.UUID) (plan.Plan, error) {

	var result plan.Plan

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ?", id)

	return result, err
}

// Retrieves a plan by its id, and locks the row
func GetPlanForUpdate(ctx context.Context, conn database.Conn, id uuid.UUID) (plan.Plan, error) {

	var result plan.Plan

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE", id)

	return result, err
}

// Retrieves a plan schedule by its id, and locks the row
func GetPlanScheduleForUpdate(ctx context.Context, conn database.Conn, id uuid.UUID) (plan.PlanSchedule, error) {

	var result plan.PlanSchedule

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE", id)

	return result, err
}

// Number of active plan subscribers. Include subscribers who have not yet started, but have not cancelled / ended either
func CountPlanActiveSubscribers(ctx context.Context, conn database.Conn, id uuid.UUID) (int, error) {

	query := "SELECT COUNT(1) FROM " + database.TableName(plan.PlanSubscriber{}) + " WHERE plan_id = ? AND cancelled_date IS NULL AND end_date IS NULL"

	var count int

	_, err := conn.QueryContext(ctx, pg.Scan(&count), query, id)
	if err != nil {

		return -1, err
	}

	return count, nil

}

// Number of active plan schedule subscribers. Include subscribers who have not yet started, but have not cancelled / ended either
func CountPlanScheduleActiveSubscribers(ctx context.Context, conn database.Conn, id uuid.UUID) (int, error) {

	query := "SELECT COUNT(1) FROM " + database.TableName(plan.PlanSubscriber{}) + " WHERE plan_schedule_id = ? AND cancelled_date IS NULL AND end_date IS NULL"

	var count int

	_, err := conn.QueryContext(ctx, pg.Scan(&count), query, id)
	if err != nil {

		return -1, err
	}

	return count, nil

}
