package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"anuario/interno/model"
	"anuario/interno/repository"
)

var AlunoRepo = &repository.AlunoRepository{}

func ListarAlunos(w http.ResponseWriter, r *http.Request) {
	alunos := AlunoRepo.ListarTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alunos)
}

func ObterAluno(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	aluno, encontrado := AlunoRepo.BuscarPorID(id)
	if !encontrado {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aluno)
}

func CriarAluno(w http.ResponseWriter, r *http.Request) {
	var novo model.Aluno
	if err := json.NewDecoder(r.Body).Decode(&novo); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	alunoCriado := AlunoRepo.Adicionar(novo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alunoCriado)
}

func AtualizarAluno(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var dados model.Aluno
	if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	aluno, ok := AlunoRepo.Atualizar(id, dados)
	if !ok {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aluno)
}

func RemoverAluno(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !AlunoRepo.Remover(id) {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}