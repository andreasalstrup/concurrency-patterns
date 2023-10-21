package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool) // can hold value of any type
	go func() {
		time.Sleep(1 * time.Second)
		<-c
	}()
	start := time.Now()
	c <- true // blocks for 1 second, until other gorutine receives
	fmt.Printf("send took %v\n", time.Since(start))
}
