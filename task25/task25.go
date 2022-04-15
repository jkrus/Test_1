package main

import (
	"log"
	"time"
)

func main() {
	t1 := time.Now().UnixNano()
	sleep(5 * time.Second)
	log.Println(time.Now().UnixNano() - t1)
	t1 = time.Now().UnixNano()
	time.Sleep(5 * time.Second)
	log.Println(time.Now().UnixNano() - t1)
}

func sleep(timeSleep time.Duration) {
	ch := make(chan struct{})
	start := time.Now()
	end := start.Add(timeSleep).UnixNano()
	go func() {
		for {
			if time.Now().UnixNano() >= end {
				ch <- struct{}{}
			}
		}
	}()

	<-ch
}
