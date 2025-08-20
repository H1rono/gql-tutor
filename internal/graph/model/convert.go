package model

import (
	"github.com/google/uuid"
	"github.com/h1rono/gql-tutor/internal/service"
)

func UserFromService(u *service.User) *User {
	if u == nil {
		return nil
	}
	return &User{
		ID:   uuid.UUID(u.ID).String(),
		Name: u.Name,
	}
}

func RepoFromService(r *service.Repo) *Repository {
	if r == nil {
		return nil
	}
	return &Repository{
		ID:        uuid.UUID(r.ID).String(),
		Name:      r.Name,
		Owner:     UserFromService(&r.Owner),
		CreatedAt: r.CreatedAt,
	}
}
