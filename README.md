# Curso Go Cloud

## Exercício 1

## Material

[Golang seek Examples Tutorial](https://www.golinuxcloud.com/golang-seek/)

[fortune.go](https://github.com/bmc/fortune-go/blob/master/fortune.go)

[How to Get Started with Golang – a Developer Roadmap](https://www.freecodecamp.org/news/golang-developer-roadmap/)

[The Golang Handbook – A Beginner's Guide to Learning Go](https://www.freecodecamp.org/news/learn-golang-handbook/)

[How to Convert Golang Bytes to String?](https://hackthedeveloper.com/golang-bytes-to-string/)
Funciona também para caracteres UTF8 mas no caso, teremos dois bytes por caracter.
```go
stringValue := "Hello"
byteSlice := []byte(stringValue)
```
```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	password := "mySecretPassword"

	passwordBytes := []byte(password) // Convert password string to bytes
	hash := sha256.Sum256(passwordBytes)

	fmt.Printf("Password hash: %x\n", hash) // Output: Password hash: 2250e74c6f823de9d70c2222802cd059dc970f56ed8d41d5d22d1a6d4a2ab66f
}
```
Também usado para serialização e desserialização de dados.

