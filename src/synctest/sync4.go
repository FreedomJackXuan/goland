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
	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(time.Second)
				fmt.Println("aaaaaaaaaaa")

			}
		}()
	}
	var str string
	fmt.Scan(&str)
}
