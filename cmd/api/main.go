package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"anuario/internal/controller"
	customMiddleware "anuario/internal/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)            
	r.Use(middleware.Recoverer)         
	r.Use(customMiddleware.Logger)     

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/alunos", controller.ListarAlunos)
		r.Get("/alunos/{id}", controller.ObterAluno)
		r.Post("/alunos", controller.CriarAluno)
		r.Put("/alunos/{id}", controller.AtualizarAluno)
		r.Delete("/alunos/{id}", controller.RemoverAluno)

		r.Get("/eventos", controller.ListarEventos)
	})

	log.Println("Servidor rodando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}