package repository

import (
	modelo "anuario/interno/model"
	"testing"
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
