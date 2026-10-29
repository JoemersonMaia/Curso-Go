package main

import (
	"fmt"
	"strings"
)

func main() {
	var nome string
	fmt.Println("entre com o seu nome")
	fmt.Scanf("%s", &nome)

	if strings.Contains(nome, "calvo") {
		fmt.Println("vc deve ser parente do teo")
	} else {
		fmt.Println("adotado")
	}
}
