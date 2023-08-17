package domain

import (
	"time"
)

const (
	TournamentCreatedEvent = "V1.Tournaments.TournamentCreated"
)

type TournamentCreated struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Gamer1ID    string    `json:"gamer_1_id"`
	Gamer2ID    string    `json:"gamer_2_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (TournamentCreated) Key() string { return TournamentCreatedEvent }
