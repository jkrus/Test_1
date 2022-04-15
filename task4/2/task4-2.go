package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	log.Println("Введите количество воркеров")
	var numWorkers int

	// Считываем количество воркеров из консоли
	fmt.Scan(&numWorkers)

	// 2. Второй вариант
	// создать контекст и передавать в каждый воркер не созданную кем-то непонятную переменную,  а всем понятный контекст
	// В данном случае в воркере будет организовано прослушивание контеста
	// При поступлении сигнала о закрытии программы мы отменим все дочерние контексты
	// Отмену контекства поймает воркер и выполнит корректное завершение wg.Done

	log.Println("===== Пример 1 =====")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		chlClose := make(chan os.Signal)
		// Оршанизуем "подписку" наи сигналы операционной системы
		signal.Notify(chlClose, os.Interrupt, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)

		for {
			// организуем прослушивание сигнала закрытия
			<-chlClose
			log.Println("Получен сигнал останова от операционной системы, завершаю работу ...")
			cancel()
		}
	}()

	// Создадим канал, в который основная горутина будет слать данные.
	data := make(chan int32)
	// Создадим группу ожидания горутин.
	wg := &sync.WaitGroup{}

	// Запустим воркеры которые будут производить чтение из канала
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, data, i, wg)
	}
	// time.Sleep(15 * time.Second)

	// Теперь организуем запись в канал
	for {
		s := false
		select {
		case <-ctx.Done():
			s = true
			break
		default:
			d := rand.Int31()
			data <- d
		}
		if s {
			break
		}
	}

	wg.Wait()

	log.Println("Все воркеры остановлены")
	log.Println("Я Main. Работу завершил.")
}

func worker(ctx context.Context, data chan int32, numWorcker int, wg *sync.WaitGroup) {
	for {
		select {
		// Запустим прослушивание контекста
		// Если придёт сигнал о завершении, просто выйдем из функции
		case <-ctx.Done():
			log.Printf("Я воркер № %v. Завершаю работу...", numWorcker)
			// Здесь можно запустить функцию, которая будет делать полезную работу по завершению воркера
			log.Printf("Я воркер № %v. Работу завершил.", numWorcker)
			wg.Done()
			return

		case d := <-data:
			log.Printf("Я воркер № %v. Данные из канала %v: ", numWorcker, d)
			// Можно не использовать вывод в консоль. (закомментировать 2 строки выше этой и раскомментировать последний case)
			// Вывод в консоль достаточно долгий процесс относительно происходящих в этой программе.
			// Поэтому сигнал об окончании будет "распространяться" быстрее.
			// case <-data:
		}
	}
}
