package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// ReadCloser 1 - response.Body
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyBytes))
	fmt.Println("----------------------------------")

	// ReadCloser 2 - file
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	file.Write(bodyBytes)
	file.Seek(0, io.SeekStart)
	fileContent, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileContent))
	fmt.Println("----------------------------------")

	// ReadCloser 3 - gzip.NewWriter() and gzip.NewReader
	file.Truncate(0)
	file.Seek(0, io.SeekStart)

	gzWriter := gzip.NewWriter(file)
	_, err = gzWriter.Write([]byte("Writing to a writer..."))
	if err != nil {
		panic(err)
	}

	if err := gzWriter.Close(); err != nil {
		panic(fmt.Errorf("failed to close gzip writer: %w", err))
	}

	file.Seek(0, io.SeekStart)
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer gzipReader.Close()

	var buff bytes.Buffer
	io.Copy(&buff, gzipReader)
	fmt.Println(buff.String())
	fmt.Println("----------------------------------")
}
