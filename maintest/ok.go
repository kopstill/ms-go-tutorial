package main

import "fmt"

func send(ch chan string, message string) {
	ch <- message
}

func main() {
	// size := 3

	ch := make(chan string)

	// go send(ch, "one")
	// go send(ch, "two")
	// go send(ch, "three")
	// go send(ch, "four")

	send(ch, "one")
	fmt.Println(<-ch)
	send(ch, "two")
	fmt.Println(<-ch)
	send(ch, "three")
	fmt.Println(<-ch)
	send(ch, "four")
	fmt.Println(<-ch)

	fmt.Println("All data sent to the channel ...")

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}
