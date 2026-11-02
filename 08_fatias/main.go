package main

import "fmt"

func main() {
	var x [100]int
	fatia := x[:10]

	fmt.Printf("x= %T fatia= %T\n", x, fatia)

	fmt.Println("tamanho fatia:", len(fatia))

	y := []int{}
	fmt.Printf("y = %T\n", y)
	fmt.Println("tamanho y:", len(y))

	for i := 1; i <= 200; i++ {
		y = append(y, i)
		fmt.Println(len(y))
	}
	fmt.Println(y)
}
