package controller

import (
	"encoding/json"
	"net/http"

	"anuario/interno/repository"
)

var EventoRepo = &repository.EventoRepository{}

func ListarEventos(w http.ResponseWriter, r *http.Request) {
	eventos := EventoRepo.ListarTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventos)
}