package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/unitid"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type GetUnit struct {
	ID     uuid.UUID
	DB     *pg.DB
	Logger *zap.Logger
}

func (g GetUnit) logger() *zap.Logger {

	return g.Logger.Named("GetUnit")

}

func (g GetUnit) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetUnit) Call(ctx context.Context) (*workout.Unit, error) {

	loader := unitid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.ID)
	if err != nil {
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown unit ID")
	}

	return loaded, nil
}
