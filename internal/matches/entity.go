package matches

import (
	"sc-internacional/internal/championships"
	"time"
)

type Match struct {
	Id             string                     `json:"id,omitempty"`
	TeamHomeId     string                     `json:"team_home_id" binding:"required"`
	TeamAwayId     string                     `json:"team_away_id" binding:"required"`
	TeamHomeName   string                     `json:"team_home_name" binding:"required"`
	TeamAwayName   string                     `json:"team_away_name" binding:"required"`
	TeamHomeScore  int                        `json:"team_home_score" binding:"required"`
	TeamAwayScore  int                        `json:"team_away_score" binding:"required"`
	MatchDate      time.Time                  `json:"match_date" binding:"required"`
	ChampionshipId championships.Championship `json:"championship_id" binding:"required"`
}
