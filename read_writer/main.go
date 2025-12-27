package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// ReadWriter 1 - file
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	file.WriteString("Hello world.")
	file.Seek(0, io.SeekStart)
	contents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contents))
	fmt.Println("------------------------------------------")

	// ReadWriter 2 - bytes.Buffer
	var buff bytes.Buffer
	buff.WriteString("Hello from the other side.")
	contents, err = io.ReadAll(&buff)
	fmt.Println(string(contents))
	fmt.Println("------------------------------------------")

	// ReadWriter 3 - bufio.NewReadWriter()
	var buff2 bytes.Buffer

	nr := bufio.NewReader(&buff2)
	nw := bufio.NewWriter(&buff2)

	nrw := bufio.NewReadWriter(nr, nw)
	_, err = nrw.WriteString("Hello from this side of the world.")
	if err != nil {
		panic(err)
	}
	if err = nrw.Flush(); err != nil {
		panic(err)
	}

	contents, err = io.ReadAll(nrw)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contents))
	fmt.Println("------------------------------------------")
}
