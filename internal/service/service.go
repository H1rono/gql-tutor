package service

import "context"

type Service struct {
	r Repository
}

type Repository interface {
	UserRepository
	RepoRepository
}

func New(r Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) GetUser(ctx context.Context, id UserID) (*User, error) {
	return s.r.GetUser(ctx, id)
}

func (s *Service) GetUserByName(ctx context.Context, name string) (*User, error) {
	return s.r.GetUserByName(ctx, name)
}

func (s *Service) GetRepo(ctx context.Context, id RepoID) (*Repo, error) {
	return s.r.GetRepo(ctx, id)
}

func (s *Service) GetRepoByName(ctx context.Context, owner string, name string) (*Repo, error) {
	return s.r.GetRepoByName(ctx, owner, name)
}
