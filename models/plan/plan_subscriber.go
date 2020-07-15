package plan

import (
	"time"

	"github.com/gofrs/uuid"
)

type PlanSubscriber struct {
	tableName struct{} `sql:"plan.plan_subscriber"`

	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"created_at" pgxload:"defaultZero"`
	UpdatedAt             time.Time `json:"updated_at" pgxload:"defaultZero"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
	PlanID                uuid.UUID `json:"planID"`
	PlanScheduleID        uuid.UUID `json:"planScheduleID"`
	CustomerID            uuid.UUID `json:"customerID"`

	// When the plan becomes available for the customer. Required
	StartDate time.Time `json:"startDate"`

	// When the plan ends for the customer. De-normalized calculation of the start date + the plan schedule interval
	EndDate *time.Time `json:"endDate"`
	// The date the customer cancelled the plan. Nil if not cancelled
	CancelledDate *time.Time `json:"cancelledDate"`
}

func (t PlanSubscriber) TableName() string {
	return "plan.plan_subscriber"
}
