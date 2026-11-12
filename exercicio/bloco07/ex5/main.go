package main

import "fmt"

func troca(na float64, nb float64) (float64, float64) {
	na2 := nb
	nb2 := na
	return na2, nb2
}
func main() {
	var na float64
	var nb float64

	fmt.Println("informe o numero a")
	fmt.Scanf("%f", &na)

	fmt.Println("informe numero b")
	fmt.Scanf("%f", &nb)

	fmt.Println(troca(na, nb))

}
