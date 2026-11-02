package main

import "fmt"

func main() {
	total := 0.0

	for i := 1; i <= 4; i++ {
		var altura float64
		fmt.Println("informe altura")
		fmt.Scanf("%f", &altura)

		total = total + altura
	}
	fmt.Printf("total das altura: %.2f", total)

}
