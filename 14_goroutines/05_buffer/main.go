package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func pinger(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		c <- "ping"
		log.Println("mensagem enviada")
	}
	fmt.Println("envio de ping finalizdo")
}

func ponger(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		c <- "pong"
		log.Println("mensagem enviada")
	}
	fmt.Println("envio de pong finalizdo")
}

func printer(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range c {
		log.Println("mesangem recebida", msg)
		time.After(time.Millisecond * 250)
	}
	log.Println("todas as mensagens retirada do canal")
	wg.Done()
}

func main() {

	wgSenders := sync.WaitGroup{}
	canal := make(chan string, 50)

	wgPrinter := sync.WaitGroup{}
	wgPrinter.Add(1)
	go printer(canal, &wgPrinter)

	wgSenders.Add(1)
	go pinger(canal, &wgSenders)

	wgSenders.Add(1)
	go ponger(canal, &wgSenders)

	wgSenders.Wait()
	close(canal)

	wgPrinter.Wait()
}
