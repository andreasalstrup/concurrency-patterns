package main

import (
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 4; i++ {
		go doWork(c) // spawn 4 goroutines
	}

	for {
		v := <-c // revices data from goroutines - blocks until one of the goroutines sends a value
		println(v)
	}
}

func doWork(c chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
		c <- rand.Int() // sends data to main goroutine
	}
}
