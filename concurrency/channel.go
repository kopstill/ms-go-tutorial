package concurrency

import "fmt"

func channel() {
	ch := make(chan string, 1)
	sendToChannel(ch, "okay")
	readFromChannel(ch)
}

func sendToChannel(ch chan<- string, message string) {
	fmt.Printf("Sending: %v\n", message)
	ch <- message
}

func readFromChannel(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}
