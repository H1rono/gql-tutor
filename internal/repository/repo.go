package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/h1rono/gql-tutor/internal/ent"
	"github.com/h1rono/gql-tutor/internal/ent/repository"
	"github.com/h1rono/gql-tutor/internal/ent/user"
	"github.com/h1rono/gql-tutor/internal/service"
)

func (r *Repository) GetRepo(ctx context.Context, id service.RepoID) (*service.Repo, error) {
	repo, err := r.c.Repository.Get(ctx, uuid.UUID(id))
	if err != nil {
		return nil, err
	}
	owner, err := repo.QueryOwner().Only(ctx)
	if err != nil {
		return nil, err
	}
	return &service.Repo{
		ID: service.RepoID(repo.ID),
		Owner: service.User{
			ID:   service.UserID(owner.ID),
			Name: owner.Name,
		},
		Name:      repo.Name,
		CreatedAt: repo.CreatedAt,
	}, nil
}

func (r *Repository) GetRepoByName(
	ctx context.Context, ownerName string, name string,
) (*service.Repo, error) {
	repo, err := r.c.Repository.
		Query().
		WithOwner(func(q *ent.UserQuery) { q.Where(user.NameEQ(ownerName)) }).
		Where(repository.NameEQ(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	owner, err := repo.QueryOwner().Only(ctx)
	if err != nil {
		return nil, err
	}
	return &service.Repo{
		ID: service.RepoID(repo.ID),
		Owner: service.User{
			ID:   service.UserID(owner.ID),
			Name: owner.Name,
		},
		Name:      repo.Name,
		CreatedAt: repo.CreatedAt,
	}, nil
}

func (r *Repository) CreateRepo(ctx context.Context, ownerID service.UserID, name string) (*service.Repo, error) {
	repo, err := r.c.Repository.Create().
		SetOwnerID(uuid.UUID(ownerID)).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	owner, err := repo.QueryOwner().Only(ctx)
	if err != nil {
		return nil, err
	}
	return &service.Repo{
		ID: service.RepoID(repo.ID),
		Owner: service.User{
			ID:   service.UserID(owner.ID),
			Name: owner.Name,
		},
		Name:      repo.Name,
		CreatedAt: repo.CreatedAt,
	}, nil
}

var _ service.RepoRepository = (*Repository)(nil)
