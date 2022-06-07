package interfaces

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

func (c customWriter) Write(b []byte) (int, error) {
	var resp GitHubResponse
	json.Unmarshal(b, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(b), nil
}

func getGithubRepos() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(customWriter{}, resp.Body)
}
