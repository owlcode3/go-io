package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func writeAndReadData(rws io.ReadWriteSeeker) {
	dataBytes := make([]byte, 1024)

	_, err := rws.Write([]byte("Writing texts to read later..."))
	if err != nil {
		panic(err)
	}

	rws.Seek(0, io.SeekStart)

	_, err = io.ReadFull(rws, dataBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes))
}

func main() {
	// ReadWriteSeeker 1 - file
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()
	writeAndReadData(file)
	fmt.Println("------------------------------------------")
}
