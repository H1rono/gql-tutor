package service

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type RepoID uuid.UUID

type Repo struct {
	ID        RepoID
	Owner     User
	Name      string
	CreatedAt time.Time
}

type RepoCore struct {
	ID        RepoID
	OwnerID   UserID
	Name      string
	CreatedAt time.Time
}

type RepoRepository interface {
	GetRepo(ctx context.Context, id RepoID) (*Repo, error)
	GetRepoByName(ctx context.Context, owner string, name string) (*Repo, error)
}
