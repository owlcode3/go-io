package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func writeData(w io.Writer, name string) {
	data := fmt.Sprintf("This is my data: %s!!\n", name)
	_, err := w.Write([]byte(data))
	if err != nil {
		panic(err)
	}
}

func main() {
	// writer 1 - os.Stdout
	writeData(os.Stdout, "os.Stdout")
	fmt.Println("----------------------------------")

	// writer 2 - os.Stderr
	writeData(os.Stderr, "os.Stderr")
	fmt.Println("----------------------------------")

	// writer 3 - file
	file, err := os.Create("testing.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())
	writeData(file, "&os.File{}")
	fileBytes := make([]byte, 4096)
	file.Seek(0, io.SeekStart)
	io.ReadFull(file, fileBytes)
	fmt.Print(string(fileBytes))
	fmt.Println("----------------------------------")

	// writer 4 - bytes.Buffer
	var buff bytes.Buffer
	writeData(&buff, "bytes.Buffer")
	fmt.Print(buff.String())
	fmt.Println("----------------------------------")

	// writer 5 - strings.Builder
	var build strings.Builder
	writeData(&build, "strings.Builder")
	fmt.Print(build.String())
	fmt.Println("----------------------------------")
}
