package modelo

import (
	"pessoa-api-go/classes"
	"pessoa-api-go/enums"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PessoaModel struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	PrimeiroNome string
	SegundoNome  string
	UltimoNome   string
	Tipo         string
	DataInclusao time.Time
}

func PessoaModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newPessoaModel(p *classes.Pessoa) *PessoaModel {
	return &PessoaModel{
		PrimeiroNome: p.PrimeiroNome,
		SegundoNome:  p.SegundoNome,
		UltimoNome:   p.UltimoNome,
		Tipo:         string(p.Tipo),
		DataInclusao: p.DataInclusao}
}

func (p *PessoaModel) toPessoa() *classes.Pessoa {
	return &classes.Pessoa{
		Id:           p.Id.Hex(),
		PrimeiroNome: p.PrimeiroNome,
		SegundoNome:  p.SegundoNome,
		UltimoNome:   p.UltimoNome,
		Tipo:         stringParaTipoPessoa(p.Tipo)}
}

//Rodrigo: queria colocar isto em enums.tipopessoaenum mas nao consegui
func stringParaTipoPessoa(str string) enums.TipoPessoaEnum {
	switch str {
	case "PJ":
		return enums.PESSOA_JURIDICA
	default:
		return enums.PESSOA_FISICA
	}
}
