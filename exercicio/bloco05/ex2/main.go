package main

import (
	"fmt"
)

func main() {
	temps := [6]float64{}
	total := float64(0)

	for i := 1; i <= 6; i++ {
		var temp float64
		fmt.Println("entre com a temp")
		fmt.Scanf("%f", &temp)
		temps[i-1] = temp
	}

	newtemps := temps[1 : len(temps)-1]

	for _, i := range newtemps {
		total += float64(i)
	}

	media := total / float64(len(newtemps))

	if media > 30 {
		fmt.Printf("esta calor\n")
	} else {
		fmt.Println("ta frio\n")
	}
}
