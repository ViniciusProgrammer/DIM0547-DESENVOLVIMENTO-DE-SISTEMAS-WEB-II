package model 

type Aluno struct { 
	ID int `json:"id"`
	Nome string `json:"nome"`
	Foto string `json:"foto,omitempty"`
	Turma string `json:"turma,omitempty"`
}

type Evento struct {
	ID int `json:"id"`
	Titulo string `json:"titulo"`
	Descricao string `json:"descricao"`
	Data string `json:"data"`
	ImagemURL string `json:"imagem_url,omitempty"`
}