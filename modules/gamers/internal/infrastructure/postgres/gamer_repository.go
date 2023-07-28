package postgres

import (
	"context"
	"database/sql"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
)

type GamerRepository struct {
	db *sql.DB
}

var _ domain.GamerRepository = (*GamerRepository)(nil)

func NewGamerRepository(db *sql.DB) GamerRepository {
	return GamerRepository{db: db}
}

func (r GamerRepository) ChangeGamerState(ctx context.Context, id string, isApproved bool) error {

	const query = `UPDATE gamers SET is_approved=$2 WHERE id=$1`

	_, err := r.db.ExecContext(ctx, query, id, isApproved)

	return err
}

func (r GamerRepository) CreateGamer(ctx context.Context, id, firstName, lastName, email, password string, isApproved bool) error {

	const query = `INSERT INTO gamers (id, first_name, last_name, email, password, is_approved) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query, id, firstName, lastName, email, password, isApproved)

	return err
}

func (r GamerRepository) FindGamer(ctx context.Context, id string) (*domain.GamerModel, error) {
	//TODO implement me
	panic("implement me")
}

func (r GamerRepository) FindGamers(ctx context.Context) ([]*domain.GamerModel, error) {
	//TODO implement me
	panic("implement me")
}
