package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/carrilhoandre/webmotors-search-go/graph"
	"github.com/carrilhoandre/webmotors-search-go/graph/generated"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	r.Handle("/query", srv)
	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
