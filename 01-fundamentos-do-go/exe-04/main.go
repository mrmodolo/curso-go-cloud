package main

import (
	"fmt"
	"strings"
)

const frase = "Um dois três quatro cinco"

func contém(s string, substr string) {
	ok := strings.Contains(s, substr)
	if ok {
		fmt.Printf("\"%s\", contém o texto \"%s\"\n", s, substr)
	} else {
		fmt.Printf("\"%s\", não contém o texto \"%s\"\n", s, substr)
	}
}

func main() {
	contém(frase, "dois")
	contém(frase, "seis")
}
