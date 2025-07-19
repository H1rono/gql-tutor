package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/h1rono/gql-tutor/internal/ent"
	"github.com/h1rono/gql-tutor/internal/graph"
	"github.com/h1rono/gql-tutor/internal/repository"
	"github.com/h1rono/gql-tutor/internal/service"
	"github.com/vektah/gqlparser/v2/ast"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func getenv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return value, nil
}

func loadRepository() (*repository.Repository, error) {
	hostname, err := getenv("MYSQL_HOSTNAME")
	if err != nil {
		return nil, err
	}
	port, err := getenv("MYSQL_PORT")
	if err != nil {
		return nil, err
	}
	username, err := getenv("MYSQL_USER")
	if err != nil {
		return nil, err
	}
	password, err := getenv("MYSQL_PASSWORD")
	if err != nil {
		return nil, err
	}
	database, err := getenv("MYSQL_DATABASE")
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", username, password, hostname, port, database)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// TODO: migration

	return repository.New(client), nil
}

func main() {
	repository, err := loadRepository()
	if err != nil {
		log.Fatalf("failed to load repository: %v", err)
	}
	service := service.NewService(repository)

	resolver := graph.NewResolver(service)
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
