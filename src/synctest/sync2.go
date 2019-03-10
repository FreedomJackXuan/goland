package synctest

import (
	"sync"
	"fmt"
)

var total struct{
	sync.Mutex
	val int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		total.Lock()
		total.val += i
		total.Unlock()
	}
}

func Sync2() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.val)
}
