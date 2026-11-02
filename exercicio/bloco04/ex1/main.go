package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	fmt.Println("me informe uma palavra")
	fmt.Scanf("%s", &palavra)

	palavra = strings.ToLower(palavra)

	contador := 0

	for _, v := range palavra {
		if v == 'a' {
			contador++
		}
	}
	fmt.Printf("quantidade de a: %d", contador)

}
