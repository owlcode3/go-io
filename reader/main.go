package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func readData1(r io.Reader) {
	buf := make([]byte, 11)

	fmt.Println("start of readData1")
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}
	fmt.Println("end of readData1")

}

func readData2(r io.Reader) {
	var buf bytes.Buffer

	n, err := buf.ReadFrom(r)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("start of readData2")
	fmt.Printf("Read %d bytes: %s\n", n, buf.Bytes())
	fmt.Println("end of readData2")

}

func main() {
	// reader 1 - file
	file, err := os.Create("aa.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	_, _ = file.Write([]byte("Hello Boy, I need your confession."))

	file.Seek(0, io.SeekStart)
	readData1(file)

	file.Truncate(0)
	_, _ = file.Write([]byte("Hello World, this is my confession."))

	file.Seek(0, io.SeekStart)
	readData2(file)
	fmt.Println("------------------------------------------")

	// reader 2 - bufio.ReadWriter
	input := strings.NewReader("Hmm!!")
	output := &bytes.Buffer{} //

	rw := bufio.NewReadWriter(
		bufio.NewReader(input),
		bufio.NewWriter(output),
	)

	rw.Writer.WriteString("This is my world, okay!!!!")
	rw.Writer.Flush()

	readData1(rw)
	input.Seek(0, io.SeekStart)
	readData2(rw)
	fmt.Println(output.String())
	fmt.Println("------------------------------------------")

	// reader 3 - strings.NewReader()
	nr := strings.NewReader("Hello from the other side")
	readData1(nr)
	nr.Seek(0, io.SeekStart)
	readData2(nr)
	fmt.Println("------------------------------------------")

	// reader 4 - bytes.NewBufferString()
	data := "What do you want from me?"
	nbs := bytes.NewBufferString(data)
	readData1(nbs)

	// nbs.Seek(0, io.SeekStart) // bytes.NewBufferString() does not have a seek method

	nbs.WriteString(data)
	readData2(nbs)
	fmt.Println("------------------------------------------")

	// reader 5 - bytes.Buffer
	var buf bytes.Buffer
	data = "I can't just imagine myself not..."

	buf.WriteString(data)
	readData1(&buf)

	// buf.Seek(0, io.SeekStart) // bytes.Buffer does not have a seek method

	buf.WriteString(data)
	readData2(&buf)
	fmt.Println("------------------------------------------")

	// reader 6 - bufio.NewReader()
	sData := strings.NewReader("This is my story")
	var bnr = bufio.NewReader(sData)

	readData1(bnr)

	sData.Seek(0, io.SeekStart)
	readData2(bnr)
	fmt.Println("------------------------------------------")

	// reader 7 - bytes.Reader()
	br := bytes.NewReader([]byte("Talk talk talk talk"))

	readData1(br)

	br.Seek(0, io.SeekStart)
	readData2(bnr)
	fmt.Println("------------------------------------------")

	// reader 8 - http.Request.Body
	// reader 9 - http.Response.Body
}
