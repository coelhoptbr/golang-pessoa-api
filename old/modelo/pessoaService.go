package modelo

import (
	"fmt"
	"pessoa-api-go/classes"
	"pessoa-api-go/config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PessoaService struct {
	collection *mgo.Collection
}

func NewPessoaService(session *config.Session, dbName string, collectionName string) *PessoaService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(PessoaModelIndex())
	return &PessoaService{collection}
}

func (pSrv *PessoaService) CriarPessoa(p *classes.Pessoa) error {
	pesModel := newPessoaModel(p)
	return pSrv.collection.Insert(&pesModel)
}

func (pSrv *PessoaService) ObterPorId(id string) (*classes.Pessoa, error) {
	pesModel := PessoaModel{}
	err := pSrv.collection.Find(bson.M{"_id": id}).One(&pesModel)
	return pesModel.toPessoa(), err
}

func (pSrv *PessoaService) ObterTodas() (*[]classes.Pessoa, error) {
	pessoasModel := []PessoaModel{}

	err := pSrv.collection.Find(nil).All(pessoasModel)

	for _, elemento := range pessoasModel {
		pes := elemento.toPessoa()
		fmt.Printf(pes.PrimeiroNome)
	}

	return nil, err
}
