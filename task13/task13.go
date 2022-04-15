package main

import "log"

func main() {
	// признаюсь честно - гуглил подобную задачу случайно несколько дней назад.
	log.Println("===== Пример 1 =====")
	a := 10
	b := 20

	log.Println("a = ", a)
	log.Println("b = ", b)

	a = a + b
	b = a - b
	a = a - b

	log.Println("a = ", a)
	log.Println("b = ", b)

	log.Println("===== Пример 2 =====")

	var s = []int{1, 2}
	log.Println(s)
	s[0], s[1] = s[1], s[0]
	log.Println(s)

}
