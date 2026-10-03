package db

import (
	"back/domain/model"
	"context"
)

type InterUserRepository interface {
	Create(ctx context.Context, u *model.User) error
	ExistsEmail(ctx context.Context, e model.Email) (bool, error)
	Find(ctx context.Context, id string) (*model.User, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
