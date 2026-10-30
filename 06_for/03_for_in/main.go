package main

import "fmt"

func main() {
	nome := "teo calvo"

	for i := range nome {
		fmt.Println(i, nome[i], string(nome[i]))
	}
}
