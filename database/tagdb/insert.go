package tagdb

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/willtrking/pgxload"
)

func InsertTag(ctx context.Context, conn pgxload.PgxTxLoader, new tag.Tag) (*tag.Tag, error) {

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

	var result tag.Tag

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

func TagObject(ctx context.Context, conn pgxload.PgxTxLoader, tagID uuid.UUID, trainerOrganizationID uuid.UUID, objectID uuid.UUID, tagType tag.TagType) (*tag.Tagged, error) {

	newUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	newTag := tag.Tagged{
		ID:                    newUUID,
		TagID:                 tagID,
		TrainerOrganizationID: trainerOrganizationID,
		TaggedID:              objectID,
		TagType:               tagType,
	}

	ins := pgxload.NewStructInsert(database.TableName(newTag), newTag)

	ins = ins.WithReturning("*").WithConflict("DO NOTHING")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result tag.Tagged

	err = conn.Scanner(rows).Scan(&result)

	return &result, err

}

func ClearObjectTags(ctx context.Context, conn pgxload.PgxTxLoader, trainerOrganizationID uuid.UUID, objectID uuid.UUID, tagType tag.TagType) error {
	_, err := conn.Exec(ctx, pgxload.RebindPositional("DELETE FROM "+database.TableName(tag.Tagged{})+" WHERE trainer_organization_id = ? AND tagged_id = ? AND tag_type = ?"), trainerOrganizationID, objectID, tagType)
	return err
}

func TagWorkoutProgram(ctx context.Context, conn pgxload.PgxTxLoader, tagID uuid.UUID, trainerOrganizationID uuid.UUID, workoutProgramID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, workoutProgramID, tag.WorkoutProgramTagType)
}

func TagWorkoutCategory(ctx context.Context, conn pgxload.PgxTxLoader, tagID uuid.UUID, trainerOrganizationID uuid.UUID, workoutCategoryID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, workoutCategoryID, tag.WorkoutCategoryTagType)
}

func TagWorkout(ctx context.Context, conn pgxload.PgxTxLoader, tagID uuid.UUID, trainerOrganizationID uuid.UUID, workoutID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, workoutID, tag.WorkoutTagType)
}

func TagExercise(ctx context.Context, conn pgxload.PgxTxLoader, tagID uuid.UUID, trainerOrganizationID uuid.UUID, exerciseID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, exerciseID, tag.ExerciseTagType)
}

func ClearExerciseTags(ctx context.Context, conn pgxload.PgxTxLoader, trainerOrganizationID uuid.UUID, objectID uuid.UUID) error {
	return ClearObjectTags(ctx, conn, trainerOrganizationID, objectID, tag.ExerciseTagType)
}

func TagPrescription(ctx context.Context, conn pgxload.PgxTxLoader, tagID uuid.UUID, trainerOrganizationID uuid.UUID, prescriptionID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, prescriptionID, tag.PrescriptionTagType)
}

func TagPlan(ctx context.Context, conn pgxload.PgxTxLoader, tagID uuid.UUID, trainerOrganizationID uuid.UUID, planID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, planID, tag.PlanTagType)
}
