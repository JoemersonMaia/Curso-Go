package main

import "fmt"

func main() {
	fmt.Println("bom dia, qual o seu nome?")
	var n string
	fmt.Scanf("%s", &n)
	fmt.Printf("prazer te conhecer, %s\n", n)
}
