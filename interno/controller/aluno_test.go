package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListarAlunos(t *testing.T) {
	req, err := http.NewRequest("GET", "/alunos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(ListarAlunos)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Erro: a rota retornou status %v, mas esperava %v", status, http.StatusOK)
	}
}
