package internal

//go:generate go tool gqlgen generate
//go:generate go tool ent generate --feature sql/versioned-migration ./ent/schema
