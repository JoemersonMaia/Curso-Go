package main

import "fmt"

func main() {
	var nota1 float32
	fmt.Println("me informe 1 nota:")
	fmt.Scanf("%f", &nota1)
	var nota2 float32
	fmt.Println("me informe 2 nota:")
	fmt.Scanf("%f", &nota2)
	var nota3 float32
	fmt.Println("me informe 3 nota:")
	fmt.Scanf("%f", &nota3)

	media := (nota1 + nota2 + nota3) / 3

	if media >= 6 {
		fmt.Println("parabens")
	} else {
		fmt.Println("perdeu")
	}
}
