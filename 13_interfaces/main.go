package main

import (
	"fmt"
	"time"
)

type Personagem struct {
	Nome   string
	Raca   string
	Clase  string
	Nivel  int
	Forca  int
	Defesa int
	Vida   int
}

type Jogador struct {
	Personagem
	TempoGameplay int64
	UltimoLogin   time.Time
	Online        bool
}

func (j *Jogador) CalcForcaAtaque() int {
	return j.Forca + j.Nivel + 10
}

func (j *Jogador) CalcForcaDefesa() int {
	return j.Defesa + j.Nivel
}

type Criatura struct {
	Personagem
	RespawTemp time.Duration
}

func (c *Criatura) CalcForcaAtaque() int {
	return c.Forca
}

func (c *Criatura) CalcForcaDefesa() int {
	return c.Defesa
}

type Npc struct {
	Personagem
	Funcao string
}

type Atacante interface {
	CalcForcaAtaque() int
}

type Defensor interface {
	CalcForcaDefesa() int
}

type Lutador interface {
	CalcForcaDefesa() int
	CalcForcaAtaque() int
}

func CalculoDano(a Atacante, d Defensor) int {
	return a.CalcForcaAtaque() - d.CalcForcaDefesa()
}

func RinhadeGo(p1, p2 Lutador) {
	fmt.Println(CalculoDano(p1, p2))
	fmt.Println(CalculoDano(p2, p1))
}

func main() {
	jojo := Jogador{
		Personagem: Personagem{
			Nome:   "joemerson",
			Raca:   "humano",
			Clase:  "guerreiro",
			Nivel:  1,
			Forca:  10,
			Defesa: 5,
			Vida:   20,
		},
		TempoGameplay: 60,
		UltimoLogin:   time.Date(2025, 12, 4, 3, 10, 3, 0, &time.Location{}),
		Online:        true,
	}

	orc := Criatura{
		Personagem: Personagem{
			Nome:   "orc 1",
			Raca:   "orc",
			Clase:  "guerreiro",
			Nivel:  2,
			Forca:  5,
			Defesa: 2,
			Vida:   5,
		},
		RespawTemp: 10,
	}

	//	Mot := Npc{
	//		Personagem: Personagem{
	//			Nome:   "Mot",
	//			Raca:   "mob",
	//			Clase:  "gelatina",
	//			Nivel:  1,
	//			Forca:  1,
	//			Defesa: 0,
	//			Vida:   2,
	//		},
	//		Funcao: "lixo",
	//	}

	fmt.Println(orc)
	fmt.Println(jojo)

	fmt.Println("forca de ataque jojo", jojo.CalcForcaAtaque())
	fmt.Println("forca de ataque orc", orc.CalcForcaAtaque())

	fmt.Println("dano jojo", CalculoDano(&jojo, &orc))
	fmt.Println("dano orc", CalculoDano(&orc, &jojo))

	RinhadeGo(&jojo, &orc)
}
