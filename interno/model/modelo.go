package model 

type Aluno struct { 
	Id int `json:"id"`
	Nome string `json:"nome"`
	Idade int `json:"idade"`
	Curso string `json:"curso"`
}

type Evento struct {
	Id int `json:"id"`
	Titulo string `json:"titulo"`
	Descricao string `json:"descricao"`
	Data string `json:"data"`
	ImagemURL string `json:"imagem_url,omitempty"`
}