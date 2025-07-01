package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg      = sync.WaitGroup{}
	counter = 0
	m       = sync.RWMutex{}
)

func main() {
	runtime.GOMAXPROCS(100)
	for range 10 {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
