package main

import "fmt"

func recpecao(input string) string {
	fmt.Printf("ola %s, seja bem vindo\n", input)
	return input
}

func main() {
	var nome string
	fmt.Println("informe seu nome")
	fmt.Scanf("%s", &nome)
	recpecao(nome)
}
