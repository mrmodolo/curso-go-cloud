package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) String() string {
	return fmt.Sprintf("{Nome:%s, Idade:%d}", p.Nome, p.Idade)
}

func (p Pessoa) DentroDaFaixaEtaria(min, max int) bool {
	return (min <= p.Idade && p.Idade <= max)
}

func Escolhidos(candidatos []Pessoa, min, max int) []Pessoa {
	escolhidos := []Pessoa{}

	for _, c := range candidatos {
		if c.DentroDaFaixaEtaria(min, max) {
			escolhidos = append(escolhidos, c)
		}
	}
	return escolhidos
}

func main() {
	candidatos := []Pessoa{
		{"Marcelo", 59},
		{"João", 39},
		{"Ana", 25},
		{"Luísa", 24},
	}
	min := 18
	max := 100
	escolhidos := Escolhidos(candidatos, min, max)
	fmt.Printf("Entre %d e %d\n", min, max)
	fmt.Println(escolhidos)

	min = 30
	max = 100
	escolhidos = Escolhidos(candidatos, min, max)
	fmt.Printf("Entre %d e %d\n", min, max)
	fmt.Println(escolhidos)
}
