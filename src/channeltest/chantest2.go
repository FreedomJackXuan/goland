package channeltest

import "fmt"

func Chantest2() {
	ch1 := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("sender send number",i,"......")
			ch1 <- i
		}
		fmt.Println("sender close channel")
		close(ch1)
	}()

	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("receriver close channel")
			break
		}
		fmt.Println("receriver receiver an element:",elem)
	}
	fmt.Println("end!")
}
