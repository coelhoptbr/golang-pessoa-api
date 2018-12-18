package classes

import (
	"pessoa-api-go/enums"
	"time"
)

type Pessoa struct {
	Id           string               `json:"id"`
	PrimeiroNome string               `json:"primeiroNome"`
	SegundoNome  string               `json:"segundoNome"`
	UltimoNome   string               `json:"ultimoNome"`
	DataInclusao time.Time            `json:"dataInclusao"`
	Tipo         enums.TipoPessoaEnum `json:"tipo"`
}

type PessoaService interface {
	CriarPessoa(p *Pessoa) error
	ObterPorId(id string) (*Pessoa, error)
}
