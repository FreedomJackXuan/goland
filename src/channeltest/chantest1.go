package channeltest

import "fmt"

func Chantest1() {
	ch1 := make(chan int, 3)
	ch1 <- 2 //包括两步：复制元素值, 放置副本到通道内部,  完成这两步之前, 发起发送操作的那句代码会一直阻塞在哪里
	ch1 <- 1
	ch1 <- 3
	elme1 := <-ch1 //包含三步： 复制通道内的元素值,放置副本到接收方,删掉原值
	fmt.Println(elme1)
}