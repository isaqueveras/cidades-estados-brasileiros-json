package script

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

var dados *Lista

// LerArquivo ler os dados do arquivo em json
func LerArquivo() []byte {
	raw, err := ioutil.ReadFile("./cidades-por-estados.json")
	if err != nil {
		log.Fatal(err)
	}

	return raw
}

// GerarDadosNovaEstrutura modela os dados para uma nova estrutura de dados
func GerarDadosNovaEstrutura(dado []byte) *NovaLista {
	if err := json.Unmarshal(dado, &dados); err != nil {
		log.Fatal(err)
	}

	novo := new(NovaLista)
	novo.TotalCidades = 0
	for i := range dados.Estados {
		estadoID := strings.Split(uuid.New().String(), "-")[0]
		novo.Estados = append(novo.Estados, NovaEstruturaEstado{
			ID:    estadoID,
			Sigla: dados.Estados[i].Sigla,
			Nome:  dados.Estados[i].Nome,
		})

		for v := range dados.Estados[i].Cidades {
			novo.TotalCidades++
			novo.Estados[i].Cidades = append(novo.Estados[i].Cidades, Cidade{
				ID:       strings.Split(uuid.New().String(), "-")[0],
				EstadoID: estadoID,
				Nome:     dados.Estados[i].Cidades[v],
				Slug:     slug.Make(dados.Estados[i].Cidades[v]),
			})
		}
	}

	return novo
}

// SalvarDados salva os dados em um arquivo
func SalvarDados(path string, dados interface{}) {
	file, err := os.Create(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	raw, err := json.Marshal(&dados)
	if err != nil {
		log.Panic(err)
	}

	if _, err = file.Write(raw); err != nil {
		log.Panic(err)
	}
}
