package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var mu sync.Mutex

	contador := 0
	totalGoRoutine := 1000

	wg.Add(totalGoRoutine)

	for i := 0; i < totalGoRoutine; i++ {

		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(contador)
}
