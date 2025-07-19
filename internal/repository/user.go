package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/h1rono/gql-tutor/internal/ent/user"
	"github.com/h1rono/gql-tutor/internal/service"
)

func (r *Repository) GetUser(ctx context.Context, id service.UserID) (*service.User, error) {
	user, err := r.c.User.Get(ctx, uuid.UUID(id))
	if err != nil {
		return nil, err
	}
	return &service.User{
		ID:   service.UserID(user.ID),
		Name: user.Name,
	}, nil
}

func (r *Repository) GetUserByName(ctx context.Context, name string) (*service.User, error) {
	user, err := r.c.User.Query().Where(user.NameEQ(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &service.User{
		ID:   service.UserID(user.ID),
		Name: user.Name,
	}, nil
}

var _ service.UserRepository = (*Repository)(nil)
