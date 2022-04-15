package main

import (
	"log"
	"strconv"
	"sync"
)

var (
	lock sync.Mutex
	wg   sync.WaitGroup
)

func main() {
	var m = make(map[int]string)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		ii := i
		go func() {
			lock.Lock()
			m[ii] = strconv.Itoa(ii)
			lock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	log.Println(m)
}
