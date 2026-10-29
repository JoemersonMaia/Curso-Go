package main

import (
	"fmt"
	"strings"
)

func main() {
	var produto string
	fmt.Println("informe o produto")
	fmt.Scanf("%s", &produto)

	itens := "laranja, cerveja, miojo, carvão, picanha"
	if strings.Contains(itens, produto) {
		fmt.Println("contem")
	} else {
		fmt.Println("esse item nao tem aqui")
	}
}
