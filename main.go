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

	data := make([]byte, 8)
	var offset int64 = 0

	var s string = ""
	for {

		offset = offset + 8

		b, eof := file.Read(data)
		if eof != nil {
			break
		}

		data = data[:b]
		if i := bytes.IndexByte(data, '\n'); i > 0 {
			s += string(data[:i])
			data = data[i+1:]
			fmt.Println(s)
			s = ""
		} else {
			s = s + string(data)
		}

	}
	if len(s) != 0 {
		fmt.Println("rest of data")
		fmt.Printf("read: %s\n", s)
	}
}
