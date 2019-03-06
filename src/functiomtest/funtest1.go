package functiomtest

import "fmt"

type Printer func(contens string) (n int, err error)

func printToStd(contens string) (bytesnum int, err error) {
	return fmt.Println(contens)
}

func Funtest1() {
	var p Printer
	p = printToStd
	p("something")
}
