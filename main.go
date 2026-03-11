package main

import (
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

	for {
		i, _ := file.ReadAt(data, offset)
		fmt.Printf("read: %s\n", data)
		offset = offset + 8

		if i == 0 {
			os.Exit(0)
		}

	}
}
