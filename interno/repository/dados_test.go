package repository

import (
	modelo "anuario/interno/model"
	"testing"
)

func TestAdicionarAluno(t *testing.T) {
	repo := &AlunoRepository{}
	novoAluno := modelo.Aluno{Nome: "Teste da Silva", Turma: "2026.1"}

	// Executa a função que queremos testar
	alunoSalvo := repo.Adicionar(novoAluno)

	// Verifica se o ID foi gerado corretamente (maior que zero)
	if alunoSalvo.ID == 0 {
		t.Errorf("Erro: Esperava que o aluno recebesse um ID válido, mas recebeu %d", alunoSalvo.ID)
	}

	// Verifica se o nome foi salvo corretamente
	if alunoSalvo.Nome != "Teste da Silva" {
		t.Errorf("Erro: Esperava o nome 'Teste da Silva', mas recebeu '%s'", alunoSalvo.Nome)
	}
}
