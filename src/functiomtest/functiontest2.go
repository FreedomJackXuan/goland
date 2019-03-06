package functiomtest

import (
	"errors"
	"fmt"
)
//方案1
type operate func(x, y int) int

func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0,errors.New("invalid operation")
	}
	return op(x, y),nil
}

func add(x int, y int) int {
	return x+y
}


//方案2
type calculateFunc func(x int, y int) (int , error)

func getCalculate(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New(" operation is nil")
		}
		return op(x,y),nil
	}
}

func Functiontest2() {
	var p operate
	p = add
	a,err:=calculate(1,2,p)
	fmt.Println(a,err)

	//2
	ad := getCalculate(nil)
	result, err1 := ad(2,3)
	fmt.Println(result,err1)
}


