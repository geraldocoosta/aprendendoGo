package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)
	go send(par, impar)
	go receive(par, impar, converge)

	for v := range converge {
		fmt.Println(v)
	}
}

func send(par, impar chan int) {
	x := 100
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			par <- i
			continue
		}
		impar <- i
	}
	close(par)
	close(impar)
}

func receive(par, impar, converge chan int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range par {
			converge <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range impar {
			converge <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(converge)
}
