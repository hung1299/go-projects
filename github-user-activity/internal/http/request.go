package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Get(url string)[]byte  {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error when call api to", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when read all body", err)
		os.Exit(1)
	}
	return body
}