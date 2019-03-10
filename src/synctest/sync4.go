package synctest

import (
	"fmt"
	"sync/atomic"
	"time"
)

func Sync4() {
	var config atomic.Value
	////var t x
	//fmt.Println("aaaaaaa")
	config.Store(make(chan interface{}))
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(make(chan interface{}))
			fmt.Println("aaaaaaaaaaa")
		}
	}()
}
