package teams

import "time"

type Team struct {
	Id             string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name           string    `json:"name" binding:"required"`
	FullName       string    `json:"fullName" binding:"required"`
	Website        string    `json:"website" binding:"required"`
	FoundationDate time.Time `json:"foundationDate" binding:"required"`
}

func (t *Team) isEmpty() bool {
	return t.Id == "" && t.Name == "" && t.FullName == "" && t.Website == "" && t.FoundationDate == time.Time{}
}
