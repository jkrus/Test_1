package main

import "fmt"

func main() {
	var (
		a int
		b string
		c bool
		d chan int
		s = []interface{}{a, b, c, d}
	)

	for _, val := range s {
		switch val.(type) {
		case int:
			fmt.Printf("I am variable of %T type \n", val.(int))
		case string:
			fmt.Printf("I am variable of %T type \n", val.(string))
		case bool:
			fmt.Printf("I am variable of %T type \n", val.(bool))
		case chan int:
			fmt.Printf("I am variable of %T type \n", val.(chan int))
		}
	}
}
