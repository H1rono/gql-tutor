package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.76

import (
	"context"
	"fmt"

	"github.com/h1rono/gql-tutor/internal/graph/model"
)

// AddProjectV2ItemByID is the resolver for the addProjectV2ItemById field.
func (r *mutationResolver) AddProjectV2ItemByID(ctx context.Context, input model.AddProjectV2ItemByIDInput) (*model.AddProjectV2ItemByIDPayload, error) {
	panic(fmt.Errorf("not implemented: AddProjectV2ItemByID - addProjectV2ItemById"))
}

// Repository is the resolver for the repository field.
func (r *queryResolver) Repository(ctx context.Context, name string, owner string) (*model.Repository, error) {
	panic(fmt.Errorf("not implemented: Repository - repository"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	u, err := r.s.GetUserByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by name: %w", err)
	}
	return model.UserFromService(u), nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
