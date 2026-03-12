package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}

	lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		var s string = ""
		for {
			data := make([]byte, 8)
			b, eof := f.Read(data)
			if eof != nil {
				break
			}

			data = data[:b]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				s += string(data[:i])
				out <- s
				data = data[i+1:]
				s = ""
			}
			s += string(data)

		}
		if len(s) != 0 {
			out <- s
		}

	}()

	return out
}
