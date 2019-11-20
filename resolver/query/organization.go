package query

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/organizationcall"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/validation"
)

func (r *QueryResolver) Organization(ctx context.Context, id uuid.UUID) (*trainer.Organization, error) {

	g := organizationcall.GetOrganization{
		ID: id,
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *QueryResolver) OrganizationAvailableTags(ctx context.Context, id uuid.UUID, first int, after *string) (*connections.TagConnection, error) {

	cursor, err := cursor.DeserializeCursor(after)
	if err != nil {
		return nil, err
	}

	g := organizationcall.GetOrganizationAvailableTags{
		TrainerOrganizationID: id,
		First:                 first,
		After:                 cursor,
		DB:                    r.db,
		Logger:                r.logger,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *QueryResolver) OrganizationWorkoutCategories(ctx context.Context, id uuid.UUID, first int, after *string) (*connections.WorkoutCategoryConnection, error) {

	cursor, err := cursor.DeserializeCursor(after)
	if err != nil {
		return nil, err
	}

	g := organizationcall.GetWorkoutCategories{
		TrainerOrganizationID: id,
		First:                 first,
		After:                 cursor,
		DB:                    r.db,
		Logger:                r.logger,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
