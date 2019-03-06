package channeltest

import (
	"time"
	"fmt"
)

func Chantest3() {
	intChan := make(chan int, 1)
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})

	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("channel is close")
			break
		}
		fmt.Println("the case is select")
	}
}
