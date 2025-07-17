package service

import (
	"context"

	"github.com/google/uuid"
)

type UserID uuid.UUID

type User struct {
	ID   UserID `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	GetUser(ctx context.Context, id UserID) (*User, error)
	GetUserByName(ctx context.Context, name string) (*User, error)
	// TODO: Add more methods as needed
}
