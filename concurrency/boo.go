package concurrency

import "fmt"

func boo() {
	size := 2
	ch := make(chan string, size)
	go send(ch, "one")
	send(ch, "two")
	// fmt.Println(<-ch)
	go send(ch, "three")
	send(ch, "four")
	// go send(ch, "five")
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}

func send(ch chan string, message string) {
	ch <- message
}
