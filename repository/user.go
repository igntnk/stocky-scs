package repository

import (
	"context"
	"github.com/igntnk/stocky-scs/models"
)

const (
	UserCollection = "user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	BlockUser(ctx context.Context, id string) (string, error)
	UnblockUser(ctx context.Context, id string) (string, error)
	UpdateUser(ctx context.Context, user *models.User) (string, error)
	GetById(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
}
