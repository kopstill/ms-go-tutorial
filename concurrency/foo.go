package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

func foo() {
	start := time.Now()

	// apis := []string{
	// apis := [6]string{
	// apis := [10]string{
	apis := [...]string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch)
	}
	fmt.Println(<-ch)

	// time.Sleep(time.Second * 20)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}
