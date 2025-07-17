package service

import "context"

type Service struct {
	u UserRepository
}

func NewService(u UserRepository) *Service {
	return &Service{
		u: u,
	}
}

func (s *Service) GetUser(ctx context.Context, id UserID) (*User, error) {
	return s.u.GetUser(ctx, id)
}

func (s *Service) GetUserByName(ctx context.Context, name string) (*User, error) {
	return s.u.GetUserByName(ctx, name)
}
