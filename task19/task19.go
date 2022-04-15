package main

import (
	"fmt"
	"log"
)

func main() {
	var str string
	fmt.Scan(&str)
	log.Println(str)
	log.Println(str[0])

	runes := make([]rune, len(str))
	runes = []rune(str)
	lastIdx := len(runes) - 1
	for idx, _ := range runes {
		if idx >= lastIdx {
			break
		}
		runes[idx], runes[lastIdx] = runes[lastIdx], runes[idx]
		lastIdx--
	}

	str = string(runes)

	fmt.Println(str)
}
