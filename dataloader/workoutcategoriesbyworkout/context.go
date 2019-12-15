package workoutcategoriesbyworkout

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/dataloader"
)

const contextKey = "workoutCategoriesByWorkout"

func GetContextLoader(ctx context.Context) *WorkoutCategoriesByWorkoutLoader {
	return ctx.Value(contextKey).(*WorkoutCategoriesByWorkoutLoader)
}

func AddContextLoader(ctx *gin.Context, db *pg.DB) {
	dataloader.GinRegisterLoader(ctx, db, contextKey, NewLoader(ctx, db))
}
