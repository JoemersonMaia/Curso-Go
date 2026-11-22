package main

import (
	"log"
	"math/rand/v2"
	"sync"
)

func GeraPacote(n, max int) []int {
	values := []int{}
	for i := 1; i <= n; i++ {
		v := rand.IntN(max) + 1
		values = append(values, v)
	}
	return values
}

func Checkpac(album map[int]int, p []int) map[int]int {
	for _, f := range p {
		album[f]++
	}
	return album
}

func CompletaAlbum(n, max int) int {
	album := map[int]int{}
	nPacote := 0
	for len(album) < max {
		pacote := GeraPacote(n, max)
		album = Checkpac(album, pacote)
		nPacote++
	}
	return nPacote
}

func soma(values []int) float64 {
	tt := 0.
	for _, v := range values {
		tt += float64(v)
	}
	return tt
}

func media(values []int) float64 {
	if len(values) == 0 {
		return 0
	}
	return soma(values) / float64(len(values))
}

func Woker(qtde, n, maximo int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("worke iniciado")
	for i := 1; i <= qtde; i++ {
		c <- CompletaAlbum(n, maximo)
	}
}

func Consolidate(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	values := []int{}
	for msg := range c {
		values = append(values, msg)
	}

	mediaValues := media(values)
	log.Printf("em media e  necessario %f pacotes para completar o album", mediaValues)
}

func main() {
	n := 5
	max := 550
	N := 1000000
	nworkers := 10

	canal := make(chan int, 1000)

	wgConsolidate := sync.WaitGroup{}
	wgConsolidate.Add(1)
	go Consolidate(canal, &wgConsolidate)

	wgWorkes := sync.WaitGroup{}
	for i := 1; i <= nworkers; i++ {
		wgWorkes.Add(1)
		go Woker(int(N/nworkers), n, max, canal, &wgWorkes)
	}

	wgWorkes.Wait()
	close(canal)

	wgConsolidate.Wait()
}
