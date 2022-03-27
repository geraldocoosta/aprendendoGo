package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canalA := trabalho("A")
	canalB := trabalho("B")

	canalC := converge(canalA, canalB)

	for {
		fmt.Println(<-canalC)
	}
}

func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i*rand.Intn(1e3))
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)
	return canal
}

func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()
	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}
