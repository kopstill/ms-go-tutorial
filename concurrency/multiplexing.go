package concurrency

import (
	"fmt"
	"time"
)

func multiplexing() {
	// ch1 := make(chan string, 1)
	// ch2 := make(chan string, 1)

	ch1 := make(chan string)
	ch2 := make(chan string)

	// processing(ch1)
	// replicating(ch2)

	go processing(ch1)
	go replicating(ch2)

	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}

func processing(ch chan<- string) {
	time.Sleep(time.Second * 3)
	ch <- "processing..."
}

func replicating(ch chan<- string) {
	time.Sleep(time.Second * 1)
	ch <- "replicating..."
}
