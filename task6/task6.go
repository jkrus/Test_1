package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {

	// Для остановки горутины будем использовать:
	// канал, контекст, таймер и просто дождёмся когда горутина закончит свою работу.

	ch := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())

	// Запустим горутину, завершим её работу используя канал
	go func() {
		for {
			select {
			case <-ch:
				log.Println("Горутина 1 прекращает работу по причине закрытия канала")
				return
			default:
				time.Sleep(time.Second)
				log.Println("Горутина 1 работает")
			}
		}
	}()

	// Здесь подождём пока горутина поработает, а потом закроем канал
	// Горутина завершит работу, а мы двинемся дальше
	time.Sleep(5 * time.Second)
	close(ch)
	time.Sleep(2 * time.Second)
	ticker := time.NewTicker(5 * time.Second)

	// Запустим горутину, завершим её работу используя таймер
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Горутина 2 прекращает работу по причине истечения времени")
				return
			default:
				time.Sleep(time.Second)
				log.Println("Горутина 2 работает")
			}
		}
	}()

	// Здесь подождём пока горутина завершит работу и мы двинемся дальше
	time.Sleep(7 * time.Second)

	// Запустим горутину, завершим её работу используя контекст
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Горутина 3 прекращает работу по причине отмены контекста")
				return
			default:
				time.Sleep(time.Second)
				log.Println("Горутина 3 работает")
			}
		}
	}()

	// Здесь подождём пока горутина поработает, а потом отменим контекст
	// Горутина завершит работу, а мы двинемся дальше
	time.Sleep(5 * time.Second)
	cancel()
	// С контекстом работать также используя и его другие методы.
	// WithDeadline (), WithTimeout ()
	// Суть - отмена контекста по истечению времени.
	time.Sleep(2 * time.Second)

	wg := sync.WaitGroup{}
	// Запустим горутину, завершим её работу используя контекст
	wg.Add(1)
	go func() {
		for {
			log.Println("Горутина 4 работает")
			time.Sleep(5 * time.Second)
			break
		}
		log.Println("Горутина 4 прекращает работу по причине окончания выполняемой работы")
		wg.Done()
	}()

	wg.Wait()

	log.Println("Основная горутина прекращает работу")
}
