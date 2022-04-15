package main

import "fmt"

var (
	set1 = map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}}
	set2 = map[int]struct{}{10: {}, 11: {}, 3: {}, 4: {}, 5: {}}
	rez  = make(map[int]struct{})
)

func main() {
	for r, _ := range set1 {
		if _, ok := set2[r]; ok {
			rez[r] = struct{}{}
		}
	}

	fmt.Println(rez)
}
