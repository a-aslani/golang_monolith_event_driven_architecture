package domain

import "context"

type GamerRepository interface {
	CreateGamer(ctx context.Context, id, firstName, lastName, email, password string, isApproved bool) error
	ChangeGamerState(ctx context.Context, id string, isApproved bool) error
	FindGamer(ctx context.Context, id string) (*GamerDTO, error)
	FindGamers(ctx context.Context) ([]*GamerDTO, error)
	UpdateGamerGem(ctx context.Context, id string, amount int) error
}

type GamerDTO struct {
	ID         string
	FirstName  string
	LastName   string
	Email      string
	IsApproved bool
}
