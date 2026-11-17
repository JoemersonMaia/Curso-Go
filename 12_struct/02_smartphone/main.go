package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint
	Peso   float64
	Altura float64
	Cpf    string
}

type Smartphone struct {
	Marca         string
	Armazenamento uint
	Cor           string
	Peso          float64
	Valor         float64
	Ligado        bool
	Propietario   Pessoa
}

func (s *Smartphone) ligar() {
	s.Ligado = true
}

func (s *Smartphone) desligar() {
	s.Ligado = false
}

type IOS struct {
	Smartphone
	os string
}

func (i *IOS) DowloadAppAppleStore() {
	fmt.Println("baixando")
}

type Android struct {
	Smartphone
	OS string
}

func (a *Android) DowloadGooglePlay() {
	if a.Ligado {
		fmt.Printf(" baixando gogogogogogo")
		return
	}
	fmt.Println("o celular esta desligado")
}

func main() {

	iphone := IOS{
		Smartphone: Smartphone{
			Marca:         "apple",
			Armazenamento: 64,
			Cor:           "azul",
			Peso:          500.0,
			Valor:         6000.0,
			Ligado:        false,
		},
		os: "ios X",
	}

	S20 := Android{
		Smartphone: Smartphone{
			Marca:         "samsumg",
			Armazenamento: 64,
			Cor:           "azul",
			Peso:          500.0,
			Valor:         6000.0,
			Ligado:        false,
		},
		OS: "android xx",
	}

	fmt.Println(iphone)
	iphone.ligar()
	fmt.Println(iphone.Ligado)

	iphone.DowloadAppAppleStore()
	fmt.Println(iphone)

	S20.ligar()
	S20.DowloadGooglePlay()
	fmt.Println(S20)
}
