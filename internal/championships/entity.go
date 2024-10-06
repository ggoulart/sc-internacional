package championships

import "time"

type Championship struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required"`
}
