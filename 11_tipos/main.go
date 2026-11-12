package main

import "fmt"

type Altura float64
type Peso float64

type Celsius float64
type Fahrenheit float64

func IMC(altura Altura, peso Peso) float64 {

	return float64(peso) / float64(altura*altura)
}

func CtoF(Celsius) Fahrenheit {
	return Fahrenheit(float64(c)*9.0/5.0 + 32.0)
}

func main() {

	alturajo := Altura(1.7)
	pesojo := Peso(72.0)

	res := float64(alturajo) + float64(pesojo)
	fmt.Println(res)

	imc := IMC(alturajo, pesojo)

	fmt.Println(imc)

	c := Celsius(30)

	fnew := CtoF(c)
	fmt.Println(fnew)

}
