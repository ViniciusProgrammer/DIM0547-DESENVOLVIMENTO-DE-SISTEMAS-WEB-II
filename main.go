package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"anuario/interno/controller"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// --- Rotas Públicas ---
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Bem-vindo à API do Memórias - Anuário Digital!"))
		})

		r.Get("/eventos", controller.ListarEventos)   // Endpoint 1
		r.Get("/alunos", controller.ListarAlunos)     // Endpoint 2
		r.Get("/alunos/{id}", controller.BuscarAluno) // Endpoint 3
	})

	// --- Rotas de Admin (Para teste, sem senha por enquanto) ---
	r.Route("/admin", func(r chi.Router) {
		r.Post("/alunos", controller.CriarAluno) // Endpoint 4
		r.Put("/alunos/{id}", controller.AtualizarAluno)
		r.Delete("/alunos/{id}", controller.DeletarAluno)
	})

	porta := ":8080"
	fmt.Printf("Servidor a iniciar na porta %s...\n", porta)
	err := http.ListenAndServe(porta, r)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
