// ### [Exercício 5: Agrupamento de Pessoas por Idade](https://backstage.stone.tech/docs/pagarme-psp/component/golang-lessons/modulo-1/02-controle-de-fluxo/#exercicio-5-agrupamento-de-pessoas-por-idade "Permanent link")
//
// **Objetivo**: Criar um sistema de agrupamento de pessoas por idade.
//
// 1. Dado uma slice de structs `Pessoa`, onde cada `Pessoa` tem uma idade.
// 2. Use um map para agrupar pessoas por idade, onde a chave é a idade e o valor é uma slice de `Pessoa`.
// 3. Imprima o agrupamento de modo que para cada idade listada, você mostre os nomes das pessoas dessa idade.

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

var nomes = []string{
	"Miguel", "Sophia",
	"Davi", "Alice", "Arthur", "Julia", "Davi", "Alice", "Arthur", "Julia",
	"Pedro", "Isabella", "Gabriel", "Manuela", "Bernardo", "Laura", "Lucas",
	"Luiza", "Matheus", "Valentina", "Rafael", "Giovanna", "Heitor", "Maria Eduarda",
	"Enzo", "Helena", "Guilherme", "Beatriz", "Nicolas", "Maria Luiza", "Lorenzo", "Lara",
	"Gustavo", "Mariana", "Felipe", "Nicole", "Samuel", "Rafaela", "João Pedro", "Heloísa",
	"Daniel", "Isadora", "Vitor", "Lívia", "Leonardo", "Maria Clara", "Henrique", "Ana Clara",
	"Theo", "Lorena", "Murilo", "Gabriela", "Eduardo", "Yasmin", "Pedro Henrique", "Isabelly",
	"Pietro", "Sarah", "Cauã", "Ana Julia", "Isaac", "Letícia", "Caio", "Ana Luiza", "Vinicius",
	"Melissa", "Benjamin", "Marina", "João", "Clara", "Lucca", "Cecília", "João Miguel", "Esther",
	"Bryan", "Emanuelly", "Joaquim", "Rebeca", "João Vitor", "Ana Beatriz", "Thiago", "Lavínia",
	"Antônio", "Vitória", "Davi Lucas", "Bianca", "Francisco", "Catarina", "Enzo Gabriel",
	"Larissa", "Bruno", "Maria Fernanda", "Emanuel", "Fernanda", "João Gabriel", "Amanda",
	"Ian", "Alícia", "Davi Luiz", "Carolina", "Rodrigo", "Agatha", "Otávio", "Gabrielly",
}

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) String() string {
	return fmt.Sprintf("%s (%d anos)", p.Nome, p.Idade)
}

func makePessoas(nomes []string) []Pessoa {
	var pessoas []Pessoa
	for _, nome := range nomes {
		pessoas = append(pessoas, Pessoa{nome, random(18, 100)})
	}
	return pessoas
}

func idadesEmOrdemCrescente(faixas map[int][]Pessoa) []int {
	keys := make([]int, 0, len(faixas))
	for key := range faixas {
		keys = append(keys, key)
	}
	// sort.Ints(keys)
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return keys
}

func agrupaPorIdade(pessoas []Pessoa) map[int][]Pessoa {
	idades := make(map[int][]Pessoa)
	for _, pessoa := range pessoas {
		idades[pessoa.Idade] = append(idades[pessoa.Idade], pessoa)
	}
	return idades
}

func main() {
	pessoas := makePessoas(nomes)
	faixasEtarias := agrupaPorIdade(pessoas)
	idades := idadesEmOrdemCrescente(faixasEtarias)
	// Em go, if é uma declaração e não uma expressão.
	plural := func(i int) string {
		if i == 1 {
			return ""
		}
		return "s"
	}
	for _, idade := range idades {
		pessoas := len(faixasEtarias[idade])
		fmt.Printf("Faixa: %d, %d pessoa%s\n", idade, pessoas, plural(pessoas))
		for _, p := range faixasEtarias[idade] {
			fmt.Printf("  - %s\n", p)
		}
	}
}
