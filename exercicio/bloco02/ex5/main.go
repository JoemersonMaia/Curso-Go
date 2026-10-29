package main

import "fmt"

func main() {
	fmt.Println("informe um numero:")
	var n int
	fmt.Scanf("%d", &n)
	dobro := n * 2
	fmt.Printf("o dobro é: %d\n", dobro)
}
