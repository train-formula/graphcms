package plan

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/types"
	"github.com/train-formula/graphcms/models/interval"
)

type PlanSchedule struct {
	tableName struct{} `sql:"plan.plan_schedule"`

	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"created_at" pgxload:"defaultZero"`
	UpdatedAt             time.Time `json:"updated_at" pgxload:"defaultZero"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
	PlanID                uuid.UUID `json:"planID"`

	Name        types.NullString `json:"plan"`
	Description types.NullString `json:"description"`

	PaymentInterval      interval.DiurnalIntervalInterval
	PaymentIntervalCount int

	PricePerInterval    int             `json:"pricePerInterval"`
	PriceMarkedDownFrom types.NullInt64 `json:"priceMarkedDownFrom"`

	DurationInterval      *interval.DiurnalIntervalInterval
	DurationIntervalCount types.NullInt64

	RegistrationAvailable bool `json:"registrationAvailable"`
	Archived              bool `json:"archived"`
}

func (t PlanSchedule) TableName() string {
	return "plan.plan_schedule"
}
