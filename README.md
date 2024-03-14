# Curso Go Cloud

## Ferramentas

[Go Watch - gow](https://github.com/mitranim/gow)
Go Watch: missing watch mode for the go command. It's invoked exactly like go,
but also watches Go files and reruns on changes.

## Perguntas e respostas

### [Why no error shows up when I compile a float and try to print with %d?](https://forum.golangbridge.org/t/why-no-error-shows-up-when-i-compile-a-float-and-try-to-print-with-d/3821)

From the fmt [package documentation](https://golang.org/pkg/fmt/):
```text
If an invalid argument is given for a verb, such as providing a string to %d,
the generated string will contain a description of the problem (…)

As far as the compiler is concerned, Printf is just some function from some librarythat
that takes a string and one or more parameters of unknown type:

func Printf(format string, a ...interface{}) (n int, err error)
Any type satisfies the empty interface, so the compiler has no reason to complain.
```
### Como funciona a função init do go?

[Understanding init in Go](https://www.digitalocean.com/community/tutorials/understanding-init-in-go)

In Go, the predefined init() function sets off a piece of code to run before any other part of your package.
This code will execute as soon as the package is imported, and can be used when you need your application tool
initialize in a specific state, such as when you have a specific configuration or set of resources with which
your application needs to start. It is also used when importing a side effect, a technique used to set the
state of a program by importing a specific package. This is often used to register one package with another 
to make sure that the program is considering the correct code for the task.

Although init() is a useful tool, it can sometimes make code difficult to read, since a hard-to-find init() 
instance will greatly affect the order in which the code is run. Because of this, it is important for 
developers who are new to Go to understand the facets of this function, so that they can make sure to 
use init() in a legible manner when writing code.

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

[Imprimindo struct](https://stackoverflow.com/questions/24512112/how-to-print-struct-variables-in-console)

fmt.Printf("%#v\n", s) imprime o tipo também.
```go
package main

    import "fmt"

    type Project struct {
        Id int64 `json:"project_id"`
        Title string `json:"title"`
        Name string `json:"name"`
    }

    func (p Project) String() string {
        return fmt.Sprintf("{Id:%d, Title:%s, Name:%s}", p.Id, p.Title, p.Name)
    }

    func main() {
        o := Project{Id: 4, Name: "hello", Title: "world"}
        fmt.Printf("%+v\n", o)
    }
```
#### Converter de C para Golang

[Data Types in C](https://www.geeksforgeeks.org/data-types-in-c/)

