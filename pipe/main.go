package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	reader, writer := io.Pipe()

	var data = map[string]any{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	}

	go func() {
		if err := json.NewEncoder(writer).Encode(data); err != nil {
			writer.CloseWithError(fmt.Errorf("writer failed to write: %w", err))
			return
		}
		writer.Close()
	}()

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", reader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var bodyData map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&bodyData); err != nil {
		panic(err)
	}

	bodyByte, err := json.MarshalIndent(bodyData, "", " ")
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintln(os.Stdout, string(bodyByte))
	if err != nil {
		panic(err)
	}
	fmt.Println("------------------------------------------")
}
