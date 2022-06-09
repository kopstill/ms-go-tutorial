package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

type ApiType string

const (
	Mazure                 ApiType = "mazure"
	Azure                  ApiType = "azure"
	Github                 ApiType = "github"
	Office                 ApiType = "office"
	Somewhereintheinternet ApiType = "somewhereintheinternet"
	Microsoft              ApiType = "microsoft"
)

type Api struct {
	apitype ApiType
	apiurl  string
}

type ApiResult struct {
	apitype ApiType
	result  string
}

func foo() {
	start := time.Now()

	// apis := []string{
	// apis := [6]string{
	// apis := [10]string{
	// apis := [...]string{
	// 	"https://management.azure.com",
	// 	"https://dev.azure.com",
	// 	"https://api.github.com",
	// 	"https://outlook.office.com/",
	// 	"https://api.somewhereintheinternet.com/",
	// 	"https://graph.microsoft.com",
	// }

	apis := []Api{
		{Mazure, "https://management.azure.com"},
		{Azure, "https://dev.azure.com"},
		{Github, "https://api.github.com"},
		Api{Office, "https://outlook.office.com/"},
		Api{Somewhereintheinternet, "https://api.somewhereintheinternet.com/"},
		Api{Microsoft, "https://graph.microsoft.com"},
	}

	ch := make(chan ApiResult, 3)
	for _, api := range apis {
		go checkAPI(api.apitype, api.apiurl, ch)
	}

	for i := 0; i < len(apis); i++ {
		// fmt.Printf("%T\n", <-ch)
		chval := <-ch
		fmt.Println(chval)
	}

	// fmt.Println(<-ch)

	// time.Sleep(time.Second * 20)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

// func checkAPI(apitype ApiType, apiurl string, ch chan ApiResult) {
func checkAPI(apitype ApiType, apiurl string, ch chan<- ApiResult) {
	_, err := http.Get(apiurl)
	if err != nil {
		ch <- ApiResult{apitype, fmt.Sprintf("ERROR: %s is down!", apiurl)}
		return
	}

	ch <- ApiResult{apitype, fmt.Sprintf("SUCCESS: %s is up and running!", apiurl)}
}
