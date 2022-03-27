package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup

	var contador int64 = 0
	totalGoRoutine := 1000

	wg.Add(totalGoRoutine)

	for i := 0; i < totalGoRoutine; i++ {

		go func() {
			atomic.AddInt64(&contador, 1)
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(contador)
}
