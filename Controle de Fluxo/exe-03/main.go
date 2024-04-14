// ### [Exercício 3: Mapeamento de Frequência de Palavras](https://backstage.stone.tech/docs/pagarme-psp/component/golang-lessons/modulo-1/02-controle-de-fluxo/#exercicio-3-mapeamento-de-frequencia-de-palavras "Permanent link")
//
// **Objetivo**: Implementar um programa que conta a frequência de cada palavra em um texto.
//
// 1. Dado uma string longa, use um map para contar quantas vezes cada palavra aparece.
// 2. Palavras são separadas por espaços, e você pode assumir que o texto é todo em minúsculas e sem pontuação.
// 3. Imprima cada palavra e sua contagem de frequência.
package main

import (
	"fmt"
	"strings"
)

func geraPalavras() []string {
	texto := `Thomas Sawyer vivia com sua tia Polly, seu meio-irmão Sid e sua prima Mary na
cidade de São Petersburgo, numa casa simples, colada ao rio Mississippi. Embora
não gostasse da escola, ele era um garoto muito esperto. Sua tia vivia beliscando-o
em REPRESÁLIA às suas travessuras. Mas ele conseguia escapar das mãos dela,
amolecendo seu coração de tia preocupada.`

	// texto = strings.ReplaceAll(texto, "\n", " ")
	// remover retorno de linha e pontuação
	r := strings.NewReplacer(
		"\n", " ",
		".", "",
		",", "",
		"!", "",
		"?", "",
	)
	texto = r.Replace(texto)
	palavras := strings.Split(texto, " ")
	return palavras
}

func contaPalavras(palavras []string) map[string]int {
	ocorrencias := make(map[string]int)
	for _, p := range palavras {
		ocorrencias[p] = ocorrencias[p] + 1
	}
	return ocorrencias
}

func main() {
	ocorrencias := contaPalavras(geraPalavras())
	for k, v := range ocorrencias {
		fmt.Printf("'%s': %d\n", k, v)
	}
}
