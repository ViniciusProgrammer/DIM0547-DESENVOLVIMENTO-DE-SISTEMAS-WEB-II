package controller

import (
	"encoding/json"
	"net/http"

	repositorio "anuario/interno/repository"
)

var EventoRepository = &repositorio.EventoRepository{}

func ListarEventos(w http.ResponseWriter, r *http.Request) {
	eventos := EventoRepository.ListarTodosEventos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventos)
}
