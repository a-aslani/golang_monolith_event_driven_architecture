package domain

import "time"

type TournamentV1 struct {
	Name        string
	Description string
	Gamer1ID    string
	Gamer2ID    string
	CreatedAt   time.Time
}

func (s TournamentV1) SnapshotName() string {
	return "tournaments.TournamentV1"
}
