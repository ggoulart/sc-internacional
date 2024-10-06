package teams

import "time"

type Team struct {
	Id             string    `json:"id,omitempty"`
	Name           string    `json:"name" binding:"required"`
	FullName       string    `json:"fullName" binding:"required"`
	Website        string    `json:"website" binding:"required"`
	FoundationDate time.Time `json:"foundationDate" binding:"required"`
}
