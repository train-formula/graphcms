package workoutcall

import (
	"context"
	"strings"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/database/trainerdb"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

type CreateWorkoutCategory struct {
	Request generated.CreateWorkoutCategory
	DB      *pg.DB
}

func (c CreateWorkoutCategory) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(c.Request.Name, "Name must not be empty", true),
		/*func() *gqlerror.Error {

			if c.request.Type == workout.UnknownBlockType {
				return gqlerror.Errorf("Must specify a valid category type")
			} else if c.request.Type == workout.GeneralBlockType {
				if c.request.RoundNumeral != nil || c.request.RoundText != nil || c.request.RoundUnitID != nil || c.request.DurationSeconds != nil {
					return gqlerror.Errorf("General workout categories may not specify round or duration data")
				}
			} else if c.request.Type == workout.RoundBlockType {
				if c.request.RoundNumeral == nil || c.request.RoundUnitID == nil {
					return gqlerror.Errorf("Round workout categories must specify round numeral and unit")
				}

				if c.request.DurationSeconds != nil {
					return gqlerror.Errorf("Round workout categories may not specify duration, use TimedRound")
				}
			} else if c.request.Type == workout.TimedRoundBlockType {
				if c.request.RoundNumeral == nil || c.request.RoundUnitID == nil {
					return gqlerror.Errorf("Timed round workout categories must specify round numeral and unit")
				}

				if c.request.DurationSeconds == nil {
					return gqlerror.Errorf("Timed round workout categories must specify duration")
				}
			}

			return nil

		},*/
	}
}

func (c CreateWorkoutCategory) Call(ctx context.Context) (*workout.WorkoutCategory, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	var finalCategory *workout.WorkoutCategory

	err = c.DB.RunInTransaction(func(t *pg.Tx) error {
		_, err = trainerdb.GetOrganization(ctx, t, c.Request.TrainerOrganizationID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Organization does not exist")
			}
			return err
		}

		err = validation.TagsAllExistForTrainer(ctx, t, c.Request.TrainerOrganizationID, c.Request.Tags)
		if err != nil {
			return err
		}

		newCategory := workout.WorkoutCategory{
			ID: newUuid,

			TrainerOrganizationID: c.Request.TrainerOrganizationID,

			Name:        c.Request.Name,
			Description: strings.TrimSpace(c.Request.Description),
			//Type:            c.request.Type,
			//RoundNumeral:    models.PtrIntToInt32(c.request.RoundNumeral),
			//RoundText:       c.request.RoundText,
			//RoundUnitID:     c.request.RoundUnitID,
			//DurationSeconds: models.PtrIntToInt32(c.request.DurationSeconds),
		}

		finalCategory, err = workoutdb.InsertWorkoutCategory(ctx, t, newCategory)
		if err != nil {
			return err
		}

		for _, tagUUID := range c.Request.Tags {

			_, err := tagdb.TagWorkoutCategory(ctx, t, tagUUID, c.Request.TrainerOrganizationID, finalCategory.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalCategory, nil
}
