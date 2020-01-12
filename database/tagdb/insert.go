package tagdb

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/tag"
)

func InsertTag(ctx context.Context, conn database.Conn, new tag.Tag) (*tag.Tag, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

func TagObject(ctx context.Context, conn database.Conn, tagID uuid.UUID, trainerOrganizationID uuid.UUID, objectID uuid.UUID, tagType tag.TagType) (*tag.Tagged, error) {

	newUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	newTag := &tag.Tagged{
		ID:                    newUUID,
		TagID:                 tagID,
		TrainerOrganizationID: trainerOrganizationID,
		TaggedID:              objectID,
		TagType:               tagType,
	}

	_, err = conn.ModelContext(ctx, newTag).Returning("*").OnConflict("DO NOTHING").Insert()
	if err != nil {
		return nil, err
	}

	return newTag, nil
}

func TagWorkoutProgram(ctx context.Context, conn database.Conn, tagID uuid.UUID, trainerOrganizationID uuid.UUID, workoutProgramID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, workoutProgramID, tag.WorkoutProgramTagType)
}

func TagWorkoutCategory(ctx context.Context, conn database.Conn, tagID uuid.UUID, trainerOrganizationID uuid.UUID, workoutCategoryID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, workoutCategoryID, tag.WorkoutCategoryTagType)
}

func TagWorkout(ctx context.Context, conn database.Conn, tagID uuid.UUID, trainerOrganizationID uuid.UUID, workoutID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, workoutID, tag.WorkoutTagType)
}

func TagExercise(ctx context.Context, conn database.Conn, tagID uuid.UUID, trainerOrganizationID uuid.UUID, exerciseID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, exerciseID, tag.ExerciseTagType)
}

func TagPrescription(ctx context.Context, conn database.Conn, tagID uuid.UUID, trainerOrganizationID uuid.UUID, prescriptionID uuid.UUID) (*tag.Tagged, error) {
	return TagObject(ctx, conn, tagID, trainerOrganizationID, prescriptionID, tag.PrescriptionTagType)
}
