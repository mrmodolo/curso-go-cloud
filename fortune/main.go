package main

import (
	"fmt"
	"fortune/strfile"
	"math/rand/v2"
	"os"
)

func main() {
	fortunes := "/usr/share/games/fortunes/fortunes"
	header, err := strfile.GetHeader(fortunes + ".dat")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	line := rand.Uint32N(header.NumStr)
	quote, err := strfile.GetString(fortunes, fortunes+".dat", line)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(quote)
	os.Exit(0)
}
