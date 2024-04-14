// ### [Exercício 4: Jogo da Forca](https://backstage.stone.tech/docs/pagarme-psp/component/golang-lessons/modulo-1/02-controle-de-fluxo/#exercicio-4-jogo-da-forca "Permanent link")
//
// **Objetivo**: Desenvolver um simples jogo da forca.
//
// 1. Escolha uma palavra secreta e represente-a por underscores (por exemplo, "_ _ _ _ _" para "casa").
// 2. Permita ao usuário adivinhar uma letra por vez. Se a letra estiver na palavra, revele sua posição.
// 3. O usuário pode errar até 5 vezes antes de perder o jogo.
// 4. Implemente o jogo utilizando loops `for` e controle de fluxo para verificar as entradas do usuário e o estado do jogo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"unicode"
)

const palavra = "modolo"

const (
	tamanhoPalavraMinimo   = 6
	tamanhoPalavraMaximo   = 12
	numeroMaximoTentativas = 5
)

func mostraPalpite(palavra, acertos string) {
	fmt.Print("[ ")
	for _, c := range palavra {
		if strings.Contains(acertos, string(c)) {
			fmt.Printf("%c ", c)
		} else {
			fmt.Printf("%s ", "_")
		}
	}
	fmt.Println("]❓")
}

func acertou(palavra, acertos string) bool {
	for _, c := range palavra {
		if !strings.Contains(acertos, string(c)) {
			return false
		}
	}
	return true
}

func prompt(message string) (rune, error) {
	var letra rune

	fmt.Print(message)
	_, err := fmt.Scanf("%c\n", &letra)
	if err != nil {
		return letra, err
	}
	if !unicode.IsLetter(letra) {
		return letra, fmt.Errorf("%c não é uma letra válida", letra)
	}
	return unicode.ToLower(letra), nil
}

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		br := bufio.NewWriter(os.Stdout)
		logger := log.New(br, "", log.Ldate)
		s := <-c
		br.Flush()
		logger.Print(s)
		br.Flush()
		os.Exit(1)
	}()
}

func main() {
	acertos := ""
	erros := ""
	tentativas := max(len(palavra), numeroMaximoTentativas)
	tentativa := 1
	for {

		if tentativa > tentativas {
			fmt.Print("Infelizmente você não conseguiu 😭")
			os.Exit(1)
		}

		fmt.Println("Adivinhe a palavra (CTRL-C para sair!)")
		fmt.Printf("Tentativa %d de %d\n", tentativa, tentativas)
		mostraPalpite(palavra, acertos)

		letra, err := prompt("Digite uma letra: ")
		if err != nil {
			fmt.Printf("%c não é uma letra válida 😡\n", letra)
			continue
		}
		s := string(letra)

		if strings.Contains(acertos, s) {
			fmt.Printf("Você já acertou a letra %c 👀\n", letra)
			continue
		}

		if strings.Contains(erros, s) {
			fmt.Printf("Você já tentou %c 😴\n", letra)
			continue
		}

		tentativa = tentativa + 1

		if strings.Contains(palavra, s) {
			acertos = acertos + s
		} else {
			erros = erros + s
			fmt.Printf("A letra %c não faz parte da palavra 😝\n", letra)
		}

		if acertou(palavra, acertos) {
			if tentativa < tentativas {
				fmt.Printf("Parabéns! Você acertou a palavra com apenas %d tentativas 👏\n", tentativa)
			} else {
				fmt.Println("Parabéns! Você acertou a palavra 🙂")
			}
			os.Exit(0)
		}
	}
}
