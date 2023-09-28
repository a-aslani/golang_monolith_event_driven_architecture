package postgres

import (
	"context"
	"database/sql"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
	"github.com/stackus/errors"
)

type StoreRepository struct {
	db *sql.DB
}

var _ domain.StoreRepository = (*StoreRepository)(nil)

func NewStoreRepository(db *sql.DB) StoreRepository {
	return StoreRepository{
		db: db,
	}
}

func (s StoreRepository) FindAll(ctx context.Context) ([]*domain.StoreDTO, error) {

	const query = `SELECT id, name, amount, price FROM store ORDER BY id DESC`

	var rows *sql.Rows
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "querying store")
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing store rows")
		}
	}(rows)

	var stores []*domain.StoreDTO

	for rows.Next() {

		d := new(domain.StoreDTO)
		err = rows.Scan(&d.ID, &d.Name, &d.Amount, &d.Price)
		if err != nil {
			return nil, errors.Wrap(err, "scanning store")
		}

		stores = append(stores, d)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "finishing store rows")
	}

	return stores, err
}

func (s StoreRepository) Find(ctx context.Context, id string) (*domain.StoreDTO, error) {

	const query = `SELECT id, name, amount, price FROM store WHERE id=$1 LIMIT 1`

	var d domain.StoreDTO

	err := s.db.QueryRowContext(ctx, query, id).Scan(&d.ID, &d.Name, &d.Amount, &d.Price)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (s StoreRepository) Insert(ctx context.Context, id string, name string, amount int, price float64) error {

	const query = `INSERT INTO store (id, name, amount, price) VALUES ($1, $2, $3, $4)`

	_, err := s.db.ExecContext(ctx, query, id, name, amount, price)

	return err
}

func (s StoreRepository) Update(ctx context.Context, id string, name string, amount int, price float64) error {

	const query = `UPDATE store SET name=$2, amount=$3, price=$4 WHERE id=$1`

	_, err := s.db.ExecContext(ctx, query, id, name, amount, price)

	return err
}

func (s StoreRepository) Remove(ctx context.Context, id string) error {

	const query = `DELETE FROM store WHERE id=$1`

	_, err := s.db.ExecContext(ctx, query, id)

	return err
}
