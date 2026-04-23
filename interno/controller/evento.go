package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	modelo "anuario/interno/model"
	repositorio "anuario/interno/repository"
)

// Instanciando os repositórios para acessar os dados
var EventoRepository = &repositorio.EventoRepository{}
var AlunoRepo = &repositorio.AlunoRepository{}

// ==========================================
// ENDPOINTS DE EVENTOS
// ==========================================

// 1. Endpoint: Listar todos os eventos
func ListarEventos(w http.ResponseWriter, r *http.Request) {
	eventos := EventoRepository.ListarTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventos)
}

// ==========================================
// ENDPOINTS DE ALUNOS
// ==========================================

// 2. Endpoint: Listar todos os alunos
func ListarAlunos(w http.ResponseWriter, r *http.Request) {
	alunos := AlunoRepo.ListarTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alunos)
}

// 3. Endpoint: Buscar aluno por ID
func BuscarAluno(w http.ResponseWriter, r *http.Request) {
	// Pega o ID da URL usando o Chi
	idParam := chi.URLParam(r, "id")

	// Converte o ID de string para inteiro
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido. Deve ser um número.", http.StatusBadRequest)
		return
	}

	// Busca no repositório
	aluno, encontrado := AlunoRepo.BuscarPorID(id)
	if !encontrado {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aluno)
}

// 4. Endpoint: Criar um novo aluno
func CriarAluno(w http.ResponseWriter, r *http.Request) {
	var novoAluno modelo.Aluno

	// Transforma o JSON recebido no corpo da requisição em uma struct Go
	err := json.NewDecoder(r.Body).Decode(&novoAluno)
	if err != nil {
		http.Error(w, "Erro ao processar os dados do aluno", http.StatusBadRequest)
		return
	}

	// Salva no repositório
	alunoCriado := AlunoRepo.Adicionar(novoAluno)

	// Retorna sucesso (Status 201 - Created) e o aluno criado com o novo ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alunoCriado)
}

// 5. Endpoint: Atualizar um aluno
func AtualizarAluno(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var alunoAtualizado modelo.Aluno
	if err := json.NewDecoder(r.Body).Decode(&alunoAtualizado); err != nil {
		http.Error(w, "Erro ao processar os novos dados", http.StatusBadRequest)
		return
	}

	aluno, sucesso := AlunoRepo.Atualizar(id, alunoAtualizado)
	if !sucesso {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aluno)
}

// 6. Endpoint: Deletar um aluno
func DeletarAluno(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	sucesso := AlunoRepo.Deletar(id)
	if !sucesso {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	// Status 204 significa "No Content" - Ação feita com sucesso, mas não há dados para retornar
	w.WriteHeader(http.StatusNoContent)
}
