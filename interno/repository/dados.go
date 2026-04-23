package repository

import (
	"sort"

	"anuario/interno/model"
)

var alunos = []model.Aluno{
	{ID: 1, Nome: "Emily Miller", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
	{ID: 2, Nome: "Francisco Matheus", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
	{ID: 3, Nome: "Vinicios David", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
	{ID: 4, Nome: "Weuler Barbosa", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
	{ID: 5, Nome: "Maria Clara", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
	{ID: 6, Nome: "João Pedro", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
	{ID: 7, Nome: "Ana Beatriz", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
	{ID: 8, Nome: "Lucas Gabriel", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
	{ID: 9, Nome: "Sofia Vitória", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
	{ID: 10, Nome: "Matheus Henrique", Foto: "https://via.placeholder.com/150", Turma: "2026.1"},
}
var proximoIDAluno = 11

var eventos = []model.Evento{
	{ID: 1, Titulo: "Aula Inaugural", Descricao: "Primeira aula do semestre", Data: "2026-02-15"},
	{ID: 2, Titulo: "Semana de Tecnologia", Descricao: "Palestras e workshops", Data: "2026-06-20"},
	{ID: 3, Titulo: "Formatura", Descricao: "Colação de grau", Data: "2029-12-10"},
	{ID: 4, Titulo: "Feira de Estágio", Descricao: "Oportunidades de estágio", Data: "2026-09-05"},
	{ID: 5, Titulo: "Conferência de Pesquisa", Descricao: "Apresentação de trabalhos", Data: "2026-11-15"},
	{ID: 6, Titulo: "Festival de Cultura", Descricao: "Celebração da diversidade", Data: "2026-10-01"},
}

type AlunoRepository struct{}

func (r *AlunoRepository) ListarTodos() []model.Aluno {
	return alunos
}

func (r *AlunoRepository) BuscarPorID(id int) (model.Aluno, bool) {
	for _, a := range alunos {
		if a.ID == id {
			return a, true
		}
	}
	return model.Aluno{}, false
}

func (r *AlunoRepository) Adicionar(novo model.Aluno) model.Aluno {
	novo.ID = proximoIDAluno
	proximoIDAluno++
	alunos = append(alunos, novo)
	return novo
}

func (r *AlunoRepository) Atualizar(id int, dados model.Aluno) (model.Aluno, bool) {
	for i, a := range alunos {
		if a.ID == id {
			dados.ID = id
			alunos[i] = dados
			return dados, true
		}
	}
	return model.Aluno{}, false
}

func (r *AlunoRepository) Remover(id int) bool {
	for i, a := range alunos {
		if a.ID == id {
			alunos = append(alunos[:i], alunos[i+1:]...)
			return true
		}
	}
	return false
}

type EventoRepository struct{}

func (r *EventoRepository) ListarTodos() []model.Evento {
	sort.Slice(eventos, func(i, j int) bool {
		return eventos[i].Data < eventos[j].Data
	})
	return eventos
}