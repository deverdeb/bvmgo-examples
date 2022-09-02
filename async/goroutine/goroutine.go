package main

import (
	"fmt"
	"time"
)

func main() {
	// execute asyncCall() function in a lightweight thread managed by the Go runtime
	go asyncCall()
	// other application treatments
	for idx := 0; idx < 10; idx++ {
		fmt.Println("Synchronous function call", idx)
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(100 * time.Millisecond)
}

func asyncCall() {
	// asynchronous application treatments
	for idx := 0; idx < 10; idx++ {
		fmt.Println("Asynchronous function call", idx)
		time.Sleep(10 * time.Millisecond)
	}
}
