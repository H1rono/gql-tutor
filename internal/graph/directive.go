package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func NewDirective() DirectiveRoot {
	return DirectiveRoot{
		IsAuthenticated: isAuthenticated,
	}
}

func isAuthenticated(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
	// TODO: Implement authentication logic
	return next(ctx)
}
