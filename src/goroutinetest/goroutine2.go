package goroutinetest

import "fmt"

func producer(c chan<-int) {
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func consumer(c <- chan int, f chan <- int) {
	for {
		if v, ok := <-c; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	f <- 1 //发送数据，通知Goroutine2函数已接受完成
}

func Goroutine2()  {
	buf:=make(chan int)
	flg := make(chan int)
	go producer(buf)
	go consumer(buf,flg)
	<-flg
}