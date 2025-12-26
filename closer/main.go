package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// closer 1 - file
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()
	fmt.Println("----------------------------------")

	// closer 2 - resp.Body
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("----------------------------------")

}
