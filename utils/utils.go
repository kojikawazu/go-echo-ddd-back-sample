package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchAPI(url string) string {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	end := time.Now()

	fmt.Println("Response from", url, ":", len(body), "Time taken:", end.Sub(start))
	return string(body)
}
