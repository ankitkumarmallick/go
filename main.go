package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage of this function is: ./http-get <URL>\n")
		os.Exit(1)
	}
	uri := args[1]
	if _, error := url.ParseRequestURI(uri); error != nil {
		fmt.Printf("Invalid URL: %s\n", error)
		os.Exit(1)
	}
	response, http_err := http.Get(uri)
	if http_err != nil {
		log.Fatal(http_err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	fmt.Printf("HTTP status code: %d\n Body: %s\n", response.StatusCode, body)

}
