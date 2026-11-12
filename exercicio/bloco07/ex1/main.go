package main

import "fmt"

func recpecao(nome string) {
	fmt.Printf("ola %s, seja bem vindo\n", nome)
}

func main() {
	var nome string
	fmt.Println("informe seu nome")
	fmt.Scanf("%s", &nome)
	recpecao(nome)
}
