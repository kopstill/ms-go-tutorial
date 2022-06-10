package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func doFibonacci(number int, ch chan string) {
	x, y := 1, 1
	for i := 0; i < number; i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("Fib(%v): %v", number, x)
}

func fibonacci() {
	start := time.Now()

	size := 15
	ch := make(chan string, size)

	for i := 1; i < size; i++ {
		go doFibonacci(i, ch)
	}

	for i := 1; i < size; i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
