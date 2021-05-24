package main

import (
	"fmt"
)

func tryRecover()  {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("This is a panic"))
	b := 0
	a := 5/b
	fmt.Println(a)
}

func main() {
	tryRecover()
}
