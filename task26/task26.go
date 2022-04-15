package main

import (
	"fmt"
	"regexp"
)

var (
	mask1 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	m     = make(map[int32]bool)
)

func init() {
	for _, m1 := range mask1 {
		m[m1] = false
	}
}

func main() {
	var str string
	fmt.Scan(&str)
	reg := regexp.MustCompile(`[0-9]+|[^a-zA-Z]+`)
	if !reg.Match([]byte(str)) {
		for _, s := range str {
			if s > 90 {
				s -= 32
			}
			if m[s] {
				fmt.Println(false)
				return
			}
			m[s] = true
		}

		fmt.Println(true)
		return
	}

	fmt.Println("Введёная строка содержит не только буквы латинского алфавита")

}
