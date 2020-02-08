package domain

import "time"

type Appointment struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	AdvisorID int64     `db:"advisor_id" json:"advisor_id" validate:"required"`
	StartTime time.Time `db:"start_datetime" json:"start_datetime" validate:"required"`
	EndTime   time.Time `db:"end_datetime" json:"end_datetime" validate:"required"`
}
