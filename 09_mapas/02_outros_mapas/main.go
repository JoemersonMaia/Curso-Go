package main

import (
	"fmt"
)

func main() {

	notas := map[string][]float64{}

	notas["teo"] = []float64{3.43, 4.324, 5.65}
	notas["nah"] = []float64{2.65, 5.6, 10}
	fmt.Println(notas)

	cursos := map[string][]string{
		"teo": {"ds", "est", "py"},
		"nah": {"arquitetura"},
	}

	cursos["lar"] = []string{"arquitetura"}
	fmt.Println(cursos)
}
