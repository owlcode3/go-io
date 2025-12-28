package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func readData(rs io.ReadSeeker) {
	dataBytes1, err := io.ReadAll(rs)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes1))

	rs.Seek(0, io.SeekStart)

	dataBytes2, err := io.ReadAll(rs)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes2))
}

func main() {
	// readSeeker 1 - file
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	file.WriteString("This is my story!!!!!!!!!!!")
	file.Seek(0, io.SeekStart)
	readData(file)
	fmt.Println("------------------------------------------")

	// readSeeker 2 - strings.NewReader()
	snr := strings.NewReader("After hours")
	readData(snr)
	fmt.Println("------------------------------------------")

	// readSeeker 3 - bytes.NewReader()
	bnr := bytes.NewReader([]byte("Walk through fire for you"))
	readData(bnr)
	fmt.Println("------------------------------------------")

	// readSeeker 4 - io.NewSectionReader()
	readerAt := strings.NewReader("calm down")
	nsr := io.NewSectionReader(readerAt, 0, 9)
	readData(nsr)
	fmt.Println("------------------------------------------")
}
