package goroutinetest

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func producer1(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out<-i * factor
	}
}

func consumer1(in <-chan int) {
	for v:=range in {
		fmt.Println(v)
	}
}

func Goroutine3(){
	ch := make(chan int ,64)
	go producer1(3, ch)
	go producer1(5,ch)
	go consumer1(ch)

	//ctrl + c退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("quit ", <-sig)
}