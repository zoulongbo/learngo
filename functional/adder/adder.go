package main

import "fmt"

func adder() func (int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder  {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}


func main() {
	//a := adder()
	//for i:=0 ; i < 10 ; i++ {
		//fmt.Println(a(i))
	//}

	b := adder2(0)

	for j:= 0; j < 5; j++  {
		var s int
		s, b = b(j)
		fmt.Println(s)
	}
}
