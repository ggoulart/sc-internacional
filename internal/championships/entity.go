package championships

import "sc-internacional/internal/teams"

type Championship struct {
	Id     string       `json:"id,omitempty"`
	Name   string       `json:"name" binding:"required"`
	Season string       `json:"season" binding:"required"`
	Teams  []teams.Team `json:"teams" binding:"required"`
}

func (c *Championship) isEmpty() bool {
	return c.Id == "" && c.Name == "" && c.Season == "" && len(c.Teams) == 0
}
