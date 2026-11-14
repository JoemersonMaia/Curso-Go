package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint
	Peso   float64
	Altura float64
	Cpf    string
}

func (p Pessoa) GetImc() float64 {
	return p.Peso / (p.Altura * p.Altura)
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

func main() {

	jo := Pessoa{
		Nome:   "joemerson maia",
		Idade:  21,
		Peso:   72.20,
		Altura: 1.68,
		Cpf:    "123-321-432-00",
	}

	celjo := Smartphone{
		Marca:         "poco",
		Armazenamento: 256,
		Cor:           "Preto",
		Peso:          150.32,
		Valor:         2200.00,
		Propietario:   jo,
	}

	fmt.Println(celjo)
	fmt.Println(celjo.Marca)

	celjo.Valor += 250.0
	fmt.Println(celjo.Valor)

	fmt.Println(celjo.Ligado)
	celjo.ligar()
	fmt.Println(celjo.Ligado)
	celjo.desligar()
	fmt.Println(celjo.Ligado)

	fmt.Println("IMC", celjo.Propietario.GetImc())
}
