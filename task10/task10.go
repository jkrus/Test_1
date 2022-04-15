package main

import (
	"fmt"
)

var (
	stepTemp = 10
	temps    = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 20, 19.9, -121, 5, 0}
)

func main() {
	set := make(map[float64][]float64)

	for _, r := range temps {
		s := float64(int(r) - int(r)%stepTemp)
		set[s] = append(set[s], r)
	}

	fmt.Println(set)
}
