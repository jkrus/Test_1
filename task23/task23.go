package main

import "log"

var (
	slice = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	idx   = 9
)

func main() {
	if idx >= len(slice) {
		log.Println("Индекс за границами массива")
		return
	}

	slice = append(slice[0:idx], slice[idx+1:]...)
	log.Println(slice)
}
