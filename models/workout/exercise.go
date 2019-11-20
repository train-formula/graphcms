package workout

type Exercise struct {
	ID                 string           `json:"id"`
	Name               string           `json:"name"`
	CategoryID         string           `json:"categoryID"`
	ExerciseText       string           `json:"exerciseText"`
	HasReps            bool             `json:"hasReps"`
	HasSets            bool             `json:"hasSets"`
	RepNumeral         *int             `json:"repNumeral"`
	RepText            *string          `json:"repText"`
	RepModifierNumeral *int             `json:"repModifierNumeral"`
	RepModifierText    *string          `json:"repModifierText"`
	SetNumeral         *int             `json:"setNumeral"`
	SetText            *string          `json:"setText"`
	Duration           *int             `json:"duration"`
	Category           *WorkoutCategory `json:"category"`
	//RepUnit            *models.Unit     `json:"repUnit"`
	//RepModifierUnit    *models.Unit     `json:"repModifierUnit"`
	//SetUnit            *models.Unit     `json:"setUnit"`
}
