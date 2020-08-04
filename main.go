package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func incrementCount() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}
}
func main() {
	wg.Add(3)
	go incrementCount()
	go incrementCount()
	go incrementCount()

	wg.Wait()

	fmt.Println(count)
}
