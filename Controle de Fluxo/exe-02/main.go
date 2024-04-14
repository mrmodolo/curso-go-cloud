package main

import (
	"fmt"
	"log"
	"strings"
)

// ### [Exercício 2: Conversor de Temperatura com Switch](https://backstage.stone.tech/docs/pagarme-psp/component/golang-lessons/modulo-1/02-controle-de-fluxo/#exercicio-2-conversor-de-temperatura-com-switch "Permanent link")
//
// **Objetivo**: Escrever um programa que converte temperaturas entre Fahrenheit, Celsius e Kelvin usando `switch`.
//
// 1. O usuário deve inserir a temperatura e a escala original (C, F, ou K).
// 2. Use `switch` para determinar a escala original e realize as conversões apropriadas para as outras duas escalas.
// 3. Imprima os resultados em todas as escalas (exceto a original).

const KelvinFactor float32 = 273.15

type Escala int8

const (
	Celsius Escala = iota
	Fahrenheight
	Kelvin
)

type Medida struct {
	Escala Escala
	Value  float32
}

func (e Escala) String() string {
	switch e {
	case Celsius:
		return "°C"
	case Fahrenheight:
		return "°F"
	default:
		return "K"
	}
}

func (m Medida) String() string {
	return fmt.Sprintf("%.2f %s", m.Value, m.Escala)
}

// F = C x 1,8 + 32
func Celsius2Fahrenheight(celsius float32) float32 {
	return celsius*1.8 + 32
}

func Convert(para Escala, medida Medida) Medida {
	switch medida.Escala {
	case Celsius:
		if para == Fahrenheight {
			return Medida{para, medida.Value*1.8 + 32}
		}
		if para == Kelvin {
			return Medida{para, medida.Value + KelvinFactor}
		}
		return medida
	case Fahrenheight:
		if para == Celsius {
			return Medida{para, (medida.Value - 32) / 1.8}
		}
		if para == Kelvin {
			return Medida{para, (medida.Value-32)/1.8 + KelvinFactor}
		}
		return medida
		// Kelvin
	default:
		if para == Celsius {
			return Medida{para, medida.Value - KelvinFactor}
		}
		if para == Fahrenheight {
			return Medida{para, (medida.Value-KelvinFactor)*1.8 + 32}
		}
		return medida
	}
}

func Prompt() Medida {
	var v float32

	fmt.Print("Digite a medida de temperatura: ")
	_, err := fmt.Scan(&v)
	if err != nil {
		// log.Fatal will print the error message and will internally call System.exit(1) so the program will terminate
		log.Fatal("valor inválido para a temperatura")
	}
	var e string
	fmt.Print("C para Celsius, F para Fahrenheight ou K para Kelvin: ")
	_, err = fmt.Scan(&e)
	if err != nil {
		log.Fatal("valor inválido para a escala")
	}
	e = strings.ToUpper(e)
	if (e != "C") && (e != "F") && (e != "K") {
		log.Fatal("valor inválido para a escala")
	}

	switch e {
	case "C":
		return Medida{Celsius, v}
	case "F":
		return Medida{Fahrenheight, v}
	default:
		return Medida{Kelvin, v}
	}
}

func Teste() {
	c := Medida{Celsius, 10.0}
	f := Medida{Fahrenheight, 42.0}
	k := Medida{Kelvin, 200.0}

	fmt.Println("De", c, "para", Convert(Kelvin, c))
	fmt.Println("De", c, "para", Convert(Fahrenheight, c))
	fmt.Println("De", c, "para", Convert(Celsius, c))

	fmt.Println("")

	fmt.Println("De", f, "para", Convert(Kelvin, f))
	fmt.Println("De", f, "para", Convert(Fahrenheight, f))
	fmt.Println("De", f, "para", Convert(Celsius, f))

	fmt.Println("")

	fmt.Println("De", k, "para", Convert(Kelvin, k))
	fmt.Println("De", k, "para", Convert(Fahrenheight, k))
	fmt.Println("De", k, "para", Convert(Celsius, k))
}

func main() {
	medida := Prompt()
	switch medida.Escala {
	case Celsius:
		fmt.Println(Convert(Fahrenheight, medida))
		fmt.Println(Convert(Kelvin, medida))
	case Fahrenheight:
		fmt.Println(Convert(Celsius, medida))
		fmt.Println(Convert(Kelvin, medida))
	default:
		fmt.Println(Convert(Celsius, medida))
		fmt.Println(Convert(Fahrenheight, medida))
	}
}
