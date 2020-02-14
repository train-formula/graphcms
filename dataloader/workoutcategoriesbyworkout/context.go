package workoutcategoriesbyworkout

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/willtrking/pgxload"
)

const contextKey = "workoutCategoriesByWorkout"

func GetContextLoader(ctx context.Context) *WorkoutCategoriesByWorkoutLoader {
	return ctx.Value(contextKey).(*WorkoutCategoriesByWorkoutLoader)
}

func AddContextLoader(ctx *gin.Context, db pgxload.PgxLoader) {
	dataloader.GinRegisterLoader(ctx, contextKey, NewLoader(ctx, db))
}
