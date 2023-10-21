package main

import "sync"

func main() {
	var wg sync.WaitGroup // waits for a collection of goroutines to finish
	for i := 0; i < 5; i++ {
		wg.Add(1) // increments the WaitGroup counter
		go func(x int) {
			sendRPC(x)
			wg.Done() // decrements the WaitGroup counter
		}(i)
	}
	wg.Wait() // blocks until the WaitGroup counter is zero
}

func sendRPC(i int) {
	println(i)
}
