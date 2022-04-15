package main

import "fmt"

var sl = []string{"cat", "cat", "dog", "cat", "tree"}
var set = make(map[string]struct{})

func main() {
	// к сожалению ничего лучше не придумал.
	for _, s := range sl {
		set[s] = struct{}{}
	}

	fmt.Println(set)
}
