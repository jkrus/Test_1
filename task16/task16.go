package main

import "log"

// Быстрая сортировка.
// Википедия гласит -
// выбери опорный элемент, всё что меньше влево,
// всё что больше или равно вправо от опорного элемента.
// Затем примени рекурсивно эту технологию, пока
// длина сортируемого массива больше 1.
// Для простоты рассчитаю опорный элемент по своему - значение среднего элемента массива.

func main() {
	s := []int{1, 2, 8, 9, 4, 7, 3, 6, 9, 8, 5, 0}
	quickSort(s)
	log.Println(s)
}

func quickSort(slice []int) {
	// Определим базовый случай
	if len(slice) > 1 {
		// Для поступившего массива оассчитаем левую и правую границы
		leftBound := 0
		rightBound := len(slice) - 1
		// Рассчитаем значение опорного элемента.
		mid := len(slice) / 2
		supportElement := slice[mid]

		for {
			if leftBound <= rightBound {
				// Будем двигаться сдева на право, пока не найдём элемент больше опорного
				for {
					log.Println(slice[leftBound])
					if slice[leftBound] < supportElement {
						leftBound++
					} else {
						break
					}
				}
				// Будем двигаться справа на лево, пока не найдём элемент меньше опорного
				for {
					log.Println(slice[rightBound])
					if slice[rightBound] > supportElement {
						rightBound--
					} else {
						break
					}
				}
				// Произведем обмен элементов, чтобы соотвтетсвовать условию - слева от опроного элемена располагаются элементы, меньше опорного
				// справа от опроного элемента располагаются элементы больше опорного
				if leftBound <= rightBound {
					slice[leftBound], slice[rightBound] = slice[rightBound], slice[leftBound]
					log.Println(slice)
					leftBound++
					rightBound--
				}
			} else {
				break
			}
		}
		// В этот момент мы имеем неотсортированный массив, но частично упордоченный относительно опонрного элемента.

		// Если правая граница > 0, значит слева есть ещё элементы, которые нужно упорядочить
		if rightBound > 0 {
			quickSort(slice[:leftBound])
		}
		// Если левая граница не дошла до последнего элемента массива, значит справа есть элементы, которые нужно упорядочить
		if leftBound < len(slice) {
			quickSort(slice[leftBound:])
		}
	}
}
