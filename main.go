package main

import "github.com/isaqueveras/cidades-estados-brasileiros-json/script"

func main() {
	gerarNovoArquivoComCidadeESlug()
}

func gerarNovoArquivoComCidadeESlug() {
	bytes := script.LerArquivo()
	dados := script.GerarDadosNovaEstrutura(bytes)
	script.SalvarDados("./novo-arquivo.json", &dados)
}
