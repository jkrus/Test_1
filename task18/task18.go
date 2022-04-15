package main

import (
	"log"
	"sync"
)

type Count struct {
	count int64
	wg    *sync.WaitGroup
	mutex sync.Mutex
}

func NewCount(wg *sync.WaitGroup) *Count {
	return &Count{wg: wg}
}

func (c *Count) Incr(delta int64) {
	c.mutex.Lock()
	c.count += delta
	c.mutex.Unlock()
	c.wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	c := NewCount(&wg)

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go c.Incr(10)
	}

	wg.Wait()

	log.Println(c.count)

}
