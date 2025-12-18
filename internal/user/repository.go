package user

import (
	"api-digital-scoring/internal/entity"
	"context"
)

type Repository interface {
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	GetById(ctx context.Context, id string) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, id string) error
}
