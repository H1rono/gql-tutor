package repository

import "github.com/h1rono/gql-tutor/internal/ent"

type Repository struct {
	c *ent.Client
}

func New(c *ent.Client) *Repository {
	return &Repository{
		c: c,
	}
}
