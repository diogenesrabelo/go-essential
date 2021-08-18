package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()
	go func() {
		escrever("Testando Goroutine!")
		waitGroup.Done()
	}()
	escrever("Sem Goroutines antes do wait")
	waitGroup.Wait()
	escrever("Sem goroutines depois do wait")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
