package plan

import (
	"time"

	"github.com/gofrs/uuid"
)

// All inventory available for a plan/schedule.
// Checking inventory for a schedule should run as follows:
// - Is there inventory available for that specific schedule? Use that and end
// - Is there inventroy available for the schedules plan? Use that and end
// - Otherwise, unlimited
// Checking inventory for a plan should run as a follows:
// - Is there inventroy available for the plan? Use that and end
// - Otherwise, unlimited
// Resolving which inventory is available should run as follows:
// - Gather all rows where PlanCustomerID is null or CustomerReleasesOn <= NOW
type PlanInventory struct {
	tableName struct{} `sql:"plan.plan_inventory"`

	ID             uuid.UUID  `json:"id"`
	CreatedAt      time.Time  `json:"created_at" pgxload:"defaultZero"`
	UpdatedAt      time.Time  `json:"updated_at" pgxload:"defaultZero"`
	PlanID         uuid.UUID  `json:"planID"`
	PlanScheduleID *uuid.UUID `json:"planScheduleID"`
	TotalInventory int        `json:"totalInventory"`
}

func (t PlanInventory) TableName() string {
	return "plan.plan_inventory"
}
