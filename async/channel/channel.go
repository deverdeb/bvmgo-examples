package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// create channel for int messages
	//channel := make(chan int) // buffer size = 0
	channel := make(chan int, 5) // specify buffer size for messages (block message sender if buffer is full)

	// a thread send 10 values
	go send(10, channel)
	// application waits all values
	receive(channel)
}

// send sends random values to channel
func send(nbValues int, channel chan int) {
	for nb := 1; nb <= nbValues; nb++ {
		value := rand.Int() % 100
		fmt.Println("Send value", value)
		channel <- value // Write value to channel
	}
	close(channel) // all messages sent, close channel
}

// receive listens channel and reads values
func receive(channel chan int) {
	// use "value := <- channel" for read only one value
	for value := range channel { // loop reads all values from channel (and exits when channel is closed)
		fmt.Println("Receive value", value)
	}
}
