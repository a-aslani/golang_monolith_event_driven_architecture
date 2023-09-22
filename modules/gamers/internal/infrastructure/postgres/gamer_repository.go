package postgres

import (
	"context"
	"database/sql"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
	"github.com/stackus/errors"
)

type GamerRepository struct {
	db *sql.DB
}

var _ domain.GamerRepository = (*GamerRepository)(nil)

func NewGamerRepository(db *sql.DB) GamerRepository {
	return GamerRepository{db: db}
}

func (r GamerRepository) UpdateGamerGem(ctx context.Context, id string, amount int) error {

	const query = `UPDATE gamers SET gem=$2 WHERE id=$1`

	_, err := r.db.ExecContext(ctx, query, id, amount)

	return err
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

func (r GamerRepository) FindGamer(ctx context.Context, id string) (*domain.GamerDTO, error) {

	const query = `SELECT first_name, last_name, email, is_approved FROM gamers WHERE id=$1 LIMIT 1`

	var gamer domain.GamerDTO

	gamer.ID = id

	err := r.db.QueryRowContext(ctx, query, id).Scan(&gamer.FirstName, &gamer.LastName, &gamer.Email, &gamer.IsApproved)
	if err != nil {
		return nil, errors.Wrap(err, "scanning gamer")
	}

	return &gamer, err
}

func (r GamerRepository) FindGamers(ctx context.Context) ([]*domain.GamerDTO, error) {

	const query = `SELECT id, first_name, last_name, email, is_approved FROM gamers`

	var rows *sql.Rows
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "querying gamers")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing gamers rows")
		}
	}(rows)

	var gamers []*domain.GamerDTO

	for rows.Next() {
		gamer := new(domain.GamerDTO)
		err := rows.Scan(&gamer.ID, &gamer.FirstName, &gamer.LastName, &gamer.Email, &gamer.IsApproved)
		if err != nil {
			return nil, errors.Wrap(err, "scanning gamer")
		}

		gamers = append(gamers, gamer)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing gamer rows")
	}

	return gamers, nil
}
