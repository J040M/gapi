package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type response map[string]interface{}

func writeJSON(w http.ResponseWriter, status int, data response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Health check
	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, response{"status": "ok"})
	})

	// Hello endpoint
	r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		writeJSON(w, http.StatusOK, response{"message": "Hello, " + name})
	})

	// Random number
	r.Get("/random", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, response{"random": rand.Intn(100)})
	})

	// Current time
	r.Get("/time", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, response{"time": time.Now().Format(time.RFC3339)})
	})

	port := ":8080"
	log.Println("Starting server on", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
