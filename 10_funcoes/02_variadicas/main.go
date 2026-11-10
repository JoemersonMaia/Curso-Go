package main

import "fmt"

func soma(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
func main() {
	var a, b, c int
	a = 10
	b = 20
	c = 40
	total := soma(a, c, b, 100, 200, 300, 400)
	fmt.Println(total)

	valor := []int{33, 3, 32, 32, 3, 43, 54, 64, 54, 32, 3}
	total = soma(valor...)
	fmt.Println(total)
}
