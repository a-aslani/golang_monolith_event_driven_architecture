package postgres

import (
	"context"
	"database/sql"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/domain"
	"time"
)

type TournamentRepository struct {
	db *sql.DB
}

var _ domain.TournamentRepository = (*TournamentRepository)(nil)

func NewTournamentRepository(db *sql.DB) TournamentRepository {
	return TournamentRepository{db: db}
}

func (r TournamentRepository) InsertTournament(ctx context.Context, id, name, description, gamer1ID, gamer2ID string, createdAt time.Time) error {

	const query = `INSERT INTO tournaments(id, name, description, gamer_1_id, gamer_2_id, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query, id, name, description, gamer1ID, gamer2ID, createdAt)

	return err
}
