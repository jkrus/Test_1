package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println("Введите количество секунд, отведённое для работы программы")
	log.Println("Каждые 200мс программа будет выводить данные в консоль")

	var t int64

	fmt.Scan(&t)
	// Создадим тикер, который просигнализирует нас по окончании отведённого времени
	ticker := time.NewTicker(time.Duration(t) * time.Second)

	// Создадим канал
	ch := make(chan int32)

	// Запустим горутину, которая будет писать в канал
	// По прошествии отведённого времени она закроет канал и прекратит свою работу
	go func() {
		stop := false
		// Создадим тикер для организации периодичной записи в канал
		// Можно и без него. С ним информативнее вроде.
		tick := time.NewTicker(time.Duration(200) * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				stop = true
			default:
				select {
				case <-tick.C:
					ch <- rand.Int31()
					tick.Reset(time.Duration(200) * time.Millisecond)
				}
			}
			if stop {
				close(ch)
				log.Println("Канал закрыт, поступление данных окончено")
				break
			}
		}
	}()

	// Организуем чтение данных из канала
	// Когда канал закроется - прекратим чтение данных и завершим работу программы

	for {
		data, ok := <-ch
		if !ok {
			log.Println("Чтение данных окончено")
			break
		}
		log.Println(data)
	}

	log.Println("Время работы программы истекло")

}
