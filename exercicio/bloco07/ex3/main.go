package main

import (
	"fmt"
	"strconv"
)

func soma(values ...float64) float64 {
	res := 0.0
	for _, v := range values {
		res += v
	}
	return res
}

func media(values ...float64) float64 {
	return soma(values...) / float64(len(values))
}

func getValues() []float64 {
	values := []float64{}
	for {
		var valorstr string
		fmt.Println("entre com um valorstr")
		fmt.Scanf("%s", &valorstr)

		if valorstr == "" {
			return values
		}

		valor, err := strconv.ParseFloat(valorstr, 64)
		if err != nil {
			fmt.Println("entre com um valorstr bom")
			continue
		}

		values = append(values, valor)
	}
}

func main() {
	valores := getValues()
	medivalor := media(valores...)

	fmt.Println("media", medivalor)
}
