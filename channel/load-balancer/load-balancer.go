package main

import "time"

type Work struct {
	x, y, z int
}

func main() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < 50; i++ {
		go worker(in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResults(out)
}

// recieves work on in channel and sends results on out channel
func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		time.Sleep(1 * time.Second) // simulate load (blocking)
		out <- w
	}
}

func sendLotsOfWork(in chan *Work) {
	print("Sending work \n")
	for i := 0; i < 1000; i++ {
		in <- &Work{i, i + i, 0}
	}
	close(in)
}

func receiveLotsOfResults(out chan *Work) {
	for w := range out {
		print("Results ", w.z, "\n")
	}
	close(out)
}
