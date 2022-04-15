package main

import (
	"log"
	"sync"
)

// Создаём срез чисел. Он проинициализируется при компилляции.
var numbers = []int{2, 4, 6, 8, 10}

func main() {
	// для запуска конкурентных вычислений будем использовать горутины.
	// поскольку горутины попадают в планировщих go, то они не выполняются в момент
	// вызова горутины. Их вызовом управляет планировщик.
	// Для ожидания завершения работы горутин нужно использовать sync.WaitGroup
	wg := sync.WaitGroup{}

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			log.Printf("%v*%v = %v", num, num, num*num)
		}(num)
	}

	wg.Wait()

}
