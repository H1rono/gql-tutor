package graph

import "github.com/h1rono/gql-tutor/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	s *service.Service
}

func NewResolver(s *service.Service) *Resolver {
	return &Resolver{
		s: s,
	}
}
