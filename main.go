package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-api-postgress/configs"
	"go-api-postgress/handlers"
	"log"
	"net/http"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	log.Printf("Server started on port %s", configs.GetSeverPort())

	r.Get("/{id}", handlers.Get)
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/all", handlers.List)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetSeverPort()), r)
}

// 3 pontos fortes fracos ameaça oportunidades
