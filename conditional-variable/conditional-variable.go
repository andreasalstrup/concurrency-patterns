package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	count := 0
	finished := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote()
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
			// must be called while holding the lock - lost wakeup problem
			cond.Broadcast() // wakes up waiting thread
		}()
	}

	mu.Lock()
	for count < 5 && finished != 10 {
		cond.Wait() // only wakes up when condition becomes true
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
	mu.Unlock()
}

func requestVote() bool {
	return true
}
