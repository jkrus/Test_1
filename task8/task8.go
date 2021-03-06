package main

import "log"

func main() {
	log.Println("===== Пример 1 =====")
	// Здесь применим операцию сдвига
	// создадим переменную, все биты == 0
	var x int64
	// Сдвинем еденицу на 0 разрядов влево
	// результат поместим в х
	x = 1 << 0
	log.Println(x)
	// Сдвинем 0 на 0 разрядов влево
	// результат поместим в х
	x = 0 << 0
	log.Println(x)

	log.Println("===== Пример 2 =====")
	// Здесь используем поразрядные операции
	// создадим переменную - все биты == 0
	var y int64
	// создадим переменную у которой установлен первый бит
	// будем использовать её в качастве маски
	var one int64 = 1
	// Применим операцию поразрядного сложения
	// которая установит первый бит в 1 в нашей испытуемой переменной
	y = one | 0
	log.Println(y)
	// Применим операцию поразрядного умножения
	// которая установит первый бит в 0 в нашей испытуемой переменной
	y = y & 0
	log.Println(y)
}
