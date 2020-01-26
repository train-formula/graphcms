// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/models"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/interval"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/models/workout"
)

type AttachUnitData struct {
	Numeral *int      `json:"numeral"`
	Text    *string   `json:"text"`
	UnitID  uuid.UUID `json:"unitID"`
}

type CreateBlockExercise struct {
	ExerciseID     uuid.UUID `json:"exerciseID"`
	PrescriptionID uuid.UUID `json:"prescriptionID"`
}

type CreateExercise struct {
	TrainerOrganizationID uuid.UUID   `json:"trainerOrganizationID"`
	Name                  string      `json:"name"`
	Description           string      `json:"description"`
	VideoURL              *string     `json:"videoURL"`
	Tags                  []uuid.UUID `json:"tags"`
}

type CreatePlan struct {
	TrainerOrganizationID uuid.UUID   `json:"trainerOrganizationID"`
	Name                  string      `json:"name"`
	Description           *string     `json:"description"`
	RegistrationAvailable bool        `json:"registrationAvailable"`
	Inventory             *int        `json:"inventory"`
	Tags                  []uuid.UUID `json:"tags"`
}

type CreatePlanSchedule struct {
	TrainerOrganizationID uuid.UUID             `json:"trainerOrganizationID"`
	PlanID                uuid.UUID             `json:"planID"`
	Name                  *string               `json:"name"`
	Description           *string               `json:"description"`
	PaymentInterval       *DiurnalIntervalInput `json:"paymentInterval"`
	PricePerInterval      int                   `json:"pricePerInterval"`
	PriceMarkedDownFrom   *int                  `json:"priceMarkedDownFrom"`
	DurationInterval      *DiurnalIntervalInput `json:"durationInterval"`
	RegistrationAvailable bool                  `json:"registrationAvailable"`
	Inventory             *int                  `json:"inventory"`
}

type CreatePrescription struct {
	TrainerOrganizationID uuid.UUID                    `json:"trainerOrganizationID"`
	Name                  string                       `json:"name"`
	PrescriptionCategory  string                       `json:"prescriptionCategory"`
	DurationSeconds       *int                         `json:"durationSeconds"`
	Sets                  []*CreatePrescriptionSetData `json:"sets"`
	Tags                  []uuid.UUID                  `json:"tags"`
}

type CreatePrescriptionSet struct {
	PrescriptionID uuid.UUID                  `json:"prescriptionID"`
	Data           *CreatePrescriptionSetData `json:"data"`
}

type CreatePrescriptionSetData struct {
	SetNumber          int             `json:"setNumber"`
	PrimaryParameter   *AttachUnitData `json:"primaryParameter"`
	SecondaryParameter *AttachUnitData `json:"secondaryParameter"`
}

type CreateTag struct {
	Tag                   string    `json:"tag"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
}

type CreateWorkout struct {
	WorkoutProgramID uuid.UUID   `json:"workoutProgramID"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	DaysFromStart    int         `json:"daysFromStart"`
	Tags             []uuid.UUID `json:"tags"`
}

type CreateWorkoutBlock struct {
	WorkoutCategoryID uuid.UUID       `json:"workoutCategoryID"`
	CategoryOrder     int             `json:"categoryOrder"`
	Round             *AttachUnitData `json:"round"`
	RoundRestDuration *int            `json:"roundRestDuration"`
	NumberOfRounds    *int            `json:"numberOfRounds"`
	DurationSeconds   *int            `json:"durationSeconds"`
}

type CreateWorkoutCategory struct {
	TrainerOrganizationID uuid.UUID   `json:"trainerOrganizationID"`
	Name                  string      `json:"name"`
	Description           string      `json:"description"`
	Tags                  []uuid.UUID `json:"tags"`
}

type CreateWorkoutProgram struct {
	TrainerOrganizationID    uuid.UUID   `json:"trainerOrganizationID"`
	Name                     string      `json:"name"`
	Description              *string     `json:"description"`
	ExactStartDate           *time.Time  `json:"exactStartDate"`
	StartsWhenCustomerStarts bool        `json:"startsWhenCustomerStarts"`
	NumberOfDays             *int        `json:"numberOfDays"`
	Tags                     []uuid.UUID `json:"tags"`
}

type DiurnalIntervalInput struct {
	Interval *interval.DiurnalIntervalInterval `json:"interval"`
	Count    int                               `json:"count"`
}

type EditExercise struct {
	ID          uuid.UUID                    `json:"id"`
	Name        *string                      `json:"name"`
	Description *string                      `json:"description"`
	VideoURL    *models.NullableStringEditor `json:"videoURL"`
}

type EditPlan struct {
	ID                    uuid.UUID                    `json:"id"`
	Name                  *string                      `json:"name"`
	Description           *models.NullableStringEditor `json:"description"`
	RegistrationAvailable *bool                        `json:"registrationAvailable"`
}

type EditPlanSchedule struct {
	ID                    uuid.UUID                    `json:"id"`
	Name                  *models.NullableStringEditor `json:"name"`
	Description           *models.NullableStringEditor `json:"description"`
	PriceMarkedDownFrom   *models.NullableIntEditor    `json:"priceMarkedDownFrom"`
	RegistrationAvailable *bool                        `json:"registrationAvailable"`
}

type EditPrescription struct {
	ID                   uuid.UUID                 `json:"id"`
	Name                 *string                   `json:"name"`
	PrescriptionCategory *string                   `json:"prescriptionCategory"`
	DurationSeconds      *models.NullableIntEditor `json:"durationSeconds"`
}

type EditPrescriptionSet struct {
	ID                 uuid.UUID               `json:"id"`
	SetNumber          *int                    `json:"setNumber"`
	PrimaryParameter   *AttachUnitData         `json:"primaryParameter"`
	SecondaryParameter *NullableAttachUnitData `json:"secondaryParameter"`
}

type EditWorkout struct {
	ID            uuid.UUID `json:"id"`
	Name          *string   `json:"name"`
	Description   *string   `json:"description"`
	DaysFromStart *int      `json:"daysFromStart"`
}

type EditWorkoutBlock struct {
	ID                uuid.UUID                 `json:"id"`
	CategoryOrder     *int                      `json:"categoryOrder"`
	Round             *NullableAttachUnitData   `json:"round"`
	RoundRestDuration *models.NullableIntEditor `json:"roundRestDuration"`
	NumberOfRounds    *models.NullableIntEditor `json:"numberOfRounds"`
	DurationSeconds   *models.NullableIntEditor `json:"durationSeconds"`
}

type EditWorkoutCategory struct {
	ID          uuid.UUID `json:"id"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
}

type ExerciseSearchRequest struct {
	TrainerOrganizationID uuid.UUID   `json:"trainerOrganizationID"`
	TagUUIDs              []uuid.UUID `json:"tagUUIDs"`
}

type ExerciseSearchResults struct {
	TagFacet *TagFacet                       `json:"tag_facet"`
	Results  *connections.ExerciseConnection `json:"results"`
}

type NullableAttachUnitData struct {
	Value *AttachUnitData `json:"value"`
}

type PlanSearchRequest struct {
	TrainerOrganizationID uuid.UUID   `json:"trainerOrganizationID"`
	TagUUIDs              []uuid.UUID `json:"tagUUIDs"`
}

type PlanSearchResults struct {
	TagFacet *TagFacet                   `json:"tag_facet"`
	Results  *connections.PlanConnection `json:"results"`
}

type PrescriptionSearchRequest struct {
	TrainerOrganizationID uuid.UUID   `json:"trainerOrganizationID"`
	TagUUIDs              []uuid.UUID `json:"tagUUIDs"`
}

type PrescriptionSearchResults struct {
	TagFacet *TagFacet                           `json:"tag_facet"`
	Results  *connections.PrescriptionConnection `json:"results"`
}

type SetWorkoutBlockExercises struct {
	WorkoutBlockID uuid.UUID              `json:"workoutBlockID"`
	BlockExercises []*CreateBlockExercise `json:"blockExercises"`
}

type SetWorkoutWorkoutCategories struct {
	WorkoutID          uuid.UUID   `json:"workoutID"`
	WorkoutCategoryIDs []uuid.UUID `json:"workoutCategoryIDs"`
}

type TagFacet struct {
	Tags []*tag.Tag `json:"tags"`
}

type WorkoutCategorySearchRequest struct {
	TrainerOrganizationID uuid.UUID   `json:"trainerOrganizationID"`
	TagUUIDs              []uuid.UUID `json:"tagUUIDs"`
}

type WorkoutCategorySearchResults struct {
	TagFacet *TagFacet                              `json:"tag_facet"`
	Results  *connections.WorkoutCategoryConnection `json:"results"`
}

type WorkoutConnection struct {
	TotalCount int              `json:"totalCount"`
	Edges      []*WorkoutEdge   `json:"edges"`
	PageInfo   *models.PageInfo `json:"pageInfo"`
}

type WorkoutEdge struct {
	Cursor uuid.UUID        `json:"cursor"`
	Node   *workout.Workout `json:"node"`
}

type WorkoutProgramSearchRequest struct {
	TrainerOrganizationID uuid.UUID   `json:"trainerOrganizationID"`
	TagUUIDs              []uuid.UUID `json:"tagUUIDs"`
}

type WorkoutProgramSearchResults struct {
	TagFacet *TagFacet                             `json:"tag_facet"`
	Results  *connections.WorkoutProgramConnection `json:"results"`
}

type ProgramLevel string

const (
	ProgramLevelBeginner     ProgramLevel = "BEGINNER"
	ProgramLevelIntermediate ProgramLevel = "INTERMEDIATE"
	ProgramLevelAdvanced     ProgramLevel = "ADVANCED"
)

var AllProgramLevel = []ProgramLevel{
	ProgramLevelBeginner,
	ProgramLevelIntermediate,
	ProgramLevelAdvanced,
}

func (e ProgramLevel) IsValid() bool {
	switch e {
	case ProgramLevelBeginner, ProgramLevelIntermediate, ProgramLevelAdvanced:
		return true
	}
	return false
}

func (e ProgramLevel) String() string {
	return string(e)
}

func (e *ProgramLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProgramLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProgramLevel", str)
	}
	return nil
}

func (e ProgramLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
