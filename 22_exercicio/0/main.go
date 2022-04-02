package main

import (
	"fmt"
	"sync"
)

func main() {
	canal := make(chan string)
	go enviaParaCanal(canal)
	recebeDoCanal(canal)

}

func recebeDoCanal(c chan string) {
	for  {
		v, ok := <- c
		if (!ok) {
			return
		}
		fmt.Println(v)
	}
}

var wg sync.WaitGroup

func enviaParaCanal(canal chan string) {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go teste(canal, i)
	}
	wg.Wait()
	close(canal)
}

func teste(c chan string, numeroFuncao int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("%d - %d", numeroFuncao, i)
	}
}
