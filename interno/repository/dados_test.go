package repository

import (
	"testing"
	modelo "anuario/interno/model"
)

func TestAdicionarAluno(t *testing.T) {
	repo := &AlunoRepository{}
	novoAluno := modelo.Aluno{Nome: "Teste da Silva", Turma: "2026.1"}

	alunoSalvo := repo.Adicionar(novoAluno)

	if alunoSalvo.ID == 0 {
		t.Errorf("Erro: Esperava que o aluno recebesse um ID válido, mas recebeu %d", alunoSalvo.ID)
	}

	if alunoSalvo.Nome != "Teste da Silva" {
		t.Errorf("Erro: Esperava o nome 'Teste da Silva', mas recebeu '%s'", alunoSalvo.Nome)
	}
}

func TestBuscarPorIDExistente(t *testing.T) {
    repo := &AlunoRepository{}
    aluno, ok := repo.BuscarPorID(1)
    if !ok || aluno.ID != 1 {
        t.Error("deveria encontrar o aluno ID 1")
    }
}

func TestBuscarPorIDInexistente(t *testing.T) {
    repo := &AlunoRepository{}
    _, ok := repo.BuscarPorID(999)
    if ok {
        t.Error("não deveria encontrar aluno com ID 999")
    }
}
