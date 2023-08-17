package domain

import (
	"context"
	"time"
)

type TournamentRepository interface {
	InsertTournament(ctx context.Context, id, name, description, gamer1ID, gamer2ID string, createdAt time.Time) error
}
