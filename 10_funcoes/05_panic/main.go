package main

import (
	"fmt"
	"strconv"
)

func validaentrada() float64 {

	defer func() {
		txt := recover()
		fmt.Println(txt)
	}()

	var input string
	fmt.Printf("entre com um floute")
	fmt.Scanf("%s", &input)

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic("entrada invalida")
	}
	return num
}
func main() {
	num := validaentrada()
	fmt.Println(num)
}
