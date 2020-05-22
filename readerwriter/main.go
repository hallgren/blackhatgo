package main

import (
	"fmt"
	"os"
)

type FooReader struct{}

func (r *FooReader) Read(p []byte) (n int, err error) {
	fmt.Print("in> ")
	return os.Stdin.Read(p)
}

type FooWriter struct{}

func (w *FooWriter) Write(p []byte) (n int, err error) {
	fmt.Print("out> ")
	return os.Stdout.Write(p)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	buffer := make([]byte, 4096)

	reader.Read(buffer)
	writer.Write(buffer)
}
