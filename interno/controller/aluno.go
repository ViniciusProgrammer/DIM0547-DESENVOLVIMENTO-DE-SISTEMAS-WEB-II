package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	modelo "anuario/internal/model"
	repositorio "anuario/internal/repository"
)

var AlunoRepository = &repositorio.AlunoRepository{}

func ListarAlunos(w http.ResponseWriter, r *http.Request) {
	alunos := AlunoRepository.ListarTodos()
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

	aluno, encontrado := AlunoRepository.BuscarPorID(id)
	if !encontrado {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aluno)
}

func CriarAluno(w http.ResponseWriter, r *http.Request) {
	var novo modelo.Aluno
	if err := json.NewDecoder(r.Body).Decode(&novo); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	alunoCriado := AlunoRepository.Adicionar(novo)

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

	var dados modelo.Aluno
	if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	aluno, ok := AlunoRepository.Atualizar(id, dados)
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

	if !AlunoRepository.Remover(id) {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}