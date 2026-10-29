package main

import (
	"fmt"
	"strings"
)

func main() {
	var nome string
	fmt.Println("informe seu nome")
	fmt.Scanf("%s", &nome)

	if strings.Contains("carvalho", nome) {
		fmt.Println("familia carvalho")
	} else if strings.Contains("ramos", nome) {
		fmt.Println("familia ramos")
	} else {
		fmt.Println("deserdado")
	}
}
