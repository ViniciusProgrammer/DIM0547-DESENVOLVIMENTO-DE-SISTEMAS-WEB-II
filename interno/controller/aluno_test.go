package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"anuario/interno/model"
)

func TestListarAlunos(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/alunos", nil)
	rec := httptest.NewRecorder()
	ListarAlunos(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("esperado 200, obtido %d", rec.Code)
	}
	if rec.Header().Get("Content-Type") != "application/json" {
		t.Error("Content-Type não é application/json")
	}
}

func TestObterAlunoExistente(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/alunos/1", nil)
	rec := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	req = req.WithContext(ctx)

	ObterAluno(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("esperado 200, obtido %d", rec.Code)
	}
	var aluno model.Aluno
	if err := json.NewDecoder(rec.Body).Decode(&aluno); err != nil {
		t.Fatal("resposta não é JSON válido")
	}
	if aluno.ID != 1 {
		t.Errorf("esperado ID 1, obtido %d", aluno.ID)
	}
}

func TestObterAlunoInexistente(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/alunos/999", nil)
	rec := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "999")
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	req = req.WithContext(ctx)

	ObterAluno(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("esperado 404, obtido %d", rec.Code)
	}
}

func TestCriarAluno(t *testing.T) {
	body := strings.NewReader(`{"nome":"Novo Aluno","foto":"url","turma":"2026.1"}`)
	req := httptest.NewRequest("POST", "/api/v1/alunos", body)
	rec := httptest.NewRecorder()

	CriarAluno(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("esperado 201, obtido %d", rec.Code)
	}
	var aluno model.Aluno
	if err := json.NewDecoder(rec.Body).Decode(&aluno); err != nil {
		t.Fatal("resposta não é JSON válido")
	}
	if aluno.ID == 0 {
		t.Error("aluno criado deve ter um ID gerado")
	}
}

func TestAtualizarAluno(t *testing.T) {
	body := strings.NewReader(`{"nome":"Atualizado","foto":"nova.jpg","turma":"2026.2"}`)
	req := httptest.NewRequest("PUT", "/api/v1/alunos/1", body)
	rec := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	req = req.WithContext(ctx)

	AtualizarAluno(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("esperado 200, obtido %d", rec.Code)
	}
	var aluno model.Aluno
	if err := json.NewDecoder(rec.Body).Decode(&aluno); err != nil {
		t.Fatal("resposta não é JSON válido")
	}
	if aluno.Nome != "Atualizado" {
		t.Errorf("esperado nome 'Atualizado', obtido '%s'", aluno.Nome)
	}
}

func TestRemoverAluno(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/api/v1/alunos/2", nil)
	rec := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "2")
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	req = req.WithContext(ctx)

	RemoverAluno(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("esperado 204, obtido %d", rec.Code)
	}
}

func TestListarEventos(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/eventos", nil)
	rec := httptest.NewRecorder()

	ListarEventos(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("esperado 200, obtido %d", rec.Code)
	}
	if rec.Header().Get("Content-Type") != "application/json" {
		t.Error("Content-Type não é application/json")
	}
	var eventos []model.Evento
	if err := json.NewDecoder(rec.Body).Decode(&eventos); err != nil {
		t.Fatal("resposta não é JSON válido")
	}
	if len(eventos) < 2 {
		t.Error("esperava pelo menos 2 eventos")
	}
}