package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	})

	log.Println("server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

var page = []byte(`
<!DOCTYPE html>
<html>
<head></head>
<h1>Go test </h1>
</html>
`)
