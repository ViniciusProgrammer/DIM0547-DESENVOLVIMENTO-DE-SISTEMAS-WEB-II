package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListarAlunos(t *testing.T) {
	// Cria uma requisição GET falsa para a nossa rota
	req, err := http.NewRequest("GET", "/alunos", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Cria um "gravador" para capturar a resposta da nossa API
	rr := httptest.NewRecorder()

	// Transforma a nossa função do controller em um Handler testável
	handler := http.HandlerFunc(ListarAlunos)

	// Simula a chamada da rota
	handler.ServeHTTP(rr, req)

	// Verifica se o Status Code retornado foi o 200 (OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Erro: a rota retornou status %v, mas esperava %v", status, http.StatusOK)
	}
}
