package service

import "context"

type Service struct {
	u UserRepository
	r RepoRepository
}

func NewService(u UserRepository, r RepoRepository) *Service {
	return &Service{
		u: u,
		r: r,
	}
}

func (s *Service) GetUser(ctx context.Context, id UserID) (*User, error) {
	return s.u.GetUser(ctx, id)
}

func (s *Service) GetUserByName(ctx context.Context, name string) (*User, error) {
	return s.u.GetUserByName(ctx, name)
}

func (s *Service) GetRepo(ctx context.Context, id RepoID) (*Repo, error) {
	return s.r.GetRepo(ctx, id)
}

func (s *Service) GetRepoByName(ctx context.Context, owner UserID, name string) (*Repo, error) {
	return s.r.GetRepoByName(ctx, owner, name)
}
