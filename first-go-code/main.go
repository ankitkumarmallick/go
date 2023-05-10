package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

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
	response, error := http.Get(uri)
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)

	if error != nil {
		log.Fatal(error)
	}
	//var words string
	if response.StatusCode != 200 {
		fmt.Printf("***********Invalid JSON object**************\nJson got: %s\n", body)
		os.Exit(1)
	}
	var words Words
	error = json.Unmarshal(body, &words)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("Words entered are: %s", strings.Join(words.Words, ","))
}
