package main

import "fmt"

func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func swapPtrs(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	var a, b int

	fmt.Println("entre com a")
	fmt.Scanf("%d", &a)

	fmt.Println("entre com b")
	fmt.Scanf("%d", &b)

	a, b = swap(a, b)
	fmt.Printf("a= %d, b= %d", a, b)

	swapPtrs(&a, &b)
	fmt.Printf("\na= %d, b= %d", a, b)
}
