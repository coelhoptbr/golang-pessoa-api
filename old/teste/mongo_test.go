package teste

import (
	"log"
	"pessoa-api-go/classes"
	"pessoa-api-go/config"
	"pessoa-api-go/enums"
	"pessoa-api-go/modelo"
	"testing"
	"time"
)

const (
	mongoUrl             = "localhost:27017"
	dbName               = "servrod"
	pessoaCollectionName = "pessoas"
)

func Test_PessoaService(t *testing.T) {
	t.Run("CriarPessoa", criarPessoa_should_insert_Pessoa_into_mongo)
}

func criarPessoa_should_insert_Pessoa_into_mongo(t *testing.T) {
	//Arrange
	session, err := config.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer session.Close()
	pessoaService := modelo.NewPessoaService(session.Copy(), dbName, pessoaCollectionName)

	testePrimeiroNome := "integration_test_primeiro_nome"
	testeSegundoNome := "integration_test_segundo_nome"
	testeUltimoNome := "integration_test_ultimo_nome"
	testeTipo := enums.PESSOA_FISICA

	pess := classes.Pessoa{
		PrimeiroNome: testePrimeiroNome,
		SegundoNome:  testeSegundoNome,
		UltimoNome:   testeUltimoNome,
		Tipo:         testeTipo,
		DataInclusao: time.Now(),
	}

	//Act
	err = pessoaService.CriarPessoa(&pess)

	//Assert
	if err != nil {
		//t.Error("Unable to create pessoa: %s", err)
		t.Error("Unable to create pessoa")
	}
	var results []classes.Pessoa
	session.GetCollection(dbName, pessoaCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got: `%i`", count)
	}
	if results[0].Id != pess.Id {
		//t.Error("Incorrect Id. Expected `%s`, Got: `%s`", testePrimeiroNome, results[0].PrimeiroNome)
		t.Error("Incorrect Id")
	}
}
