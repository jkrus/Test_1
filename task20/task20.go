package main

import (
	"fmt"
	"strings"
)

var (
	str = "1 2 3 4 5 6 7 "
	sep = " "
)

func main() {

	slice := strings.Split(str, sep)
	lastIdx := len(slice) - 1
	for idx, _ := range slice {
		if idx >= lastIdx {
			break
		}
		slice[idx], slice[lastIdx] = slice[lastIdx], slice[idx]
		lastIdx--
	}
	bytes := make([]byte, 100)
	offset := 0

	for _, val := range slice {
		offset += copy(bytes[offset:], val)
		offset += copy(bytes[offset:], sep)
	}

	fmt.Println(string(bytes[:]))
}
