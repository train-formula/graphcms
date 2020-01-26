package plan

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/models/interval"
)

type PlanSchedule struct {
	tableName struct{} `sql:"plan.plan_schedule"`

	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
	PlanID                uuid.UUID `json:"planID"`

	Name        *string `json:"plan"`
	Description *string `json:"description"`

	PaymentInterval      interval.DiurnalIntervalInterval
	PaymentIntervalCount int

	PricePerInterval    int  `json:"pricePerInterval"`
	PriceMarkedDownFrom *int `json:"priceMarkedDownFrom"`

	DurationInterval      *interval.DiurnalIntervalInterval
	DurationIntervalCount *int

	RegistrationAvailable bool `json:"registrationAvailable"`
	Archived              bool `json:"archived"`
}

func (t PlanSchedule) TableName() string {
	return "plan.plan_schedule"
}
