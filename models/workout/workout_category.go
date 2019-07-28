package workout

type WorkoutCategory struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Type      CategoryType `json:"type"`
	WorkoutID string       `json:"workoutID"`
}
