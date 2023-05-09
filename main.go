package main

import (
	"fmt"
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
		fmt.Printf("Error: %s\n", http_err)
		os.Exit(1)
	}
	fmt.Printf("Response: %v\n", response)

}
