package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan time.Time)
	ch2 := make(chan time.Time)

	go func() {
		ch1 <- time.Now()
	}()
	go func() {
		ch2 <- time.Now()
	}()

	start := time.Now()
	time.Sleep(1 * time.Second)

	select { // looks at multiple channels at once, blocks until one is ready
	case v := <-ch1:
		fmt.Println("channel 1 sends", v)
	case v := <-ch2:
		fmt.Println("channel 2 sends", v)
	default: // if no channel is ready, default case is executed (non-blocking) - optional
		fmt.Println("neither channel was ready")
	}

	fmt.Printf("select took %v\n", time.Since(start))
}
