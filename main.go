package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var s string = ""
	for {
		data := make([]byte, 8)
		b, eof := file.Read(data)
		if eof != nil {
			break
		}

		data = data[:b]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			s += string(data[:i])
			data = data[i+1:]
			fmt.Printf("read: %s\n", s)
			s = ""
		}
		s += string(data)

	}
	if len(s) != 0 {
		fmt.Printf("read: %s\n", s)
	}
}
