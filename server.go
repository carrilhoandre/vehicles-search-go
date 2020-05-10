package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/carrilhoandre/webmotors-search-go/graph"
	"github.com/carrilhoandre/webmotors-search-go/graph/generated"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	go r.Handle("/query", srv)
	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)
	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
