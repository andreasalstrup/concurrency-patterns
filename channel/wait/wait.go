package main

func main() {
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(x int) {
			sendRPC(x)
			done <- true // send data to main goroutine
		}(i)
	}
	for i := 0; i < 5; i++ {
		<-done // receive data from goroutines
	}
}

func sendRPC(i int) {
	println(i)
}
