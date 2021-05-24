package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
	"runtime"
)

func main() {
	ev, err := eval(1, 2, "xx")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("%d %s %d = %d \n", 1, "-", 2, ev)
	q, r := div(5, 3)
	fmt.Println(q, r)

	fmt.Println(apply(func(a int, b int) int { return int(math.Pow(float64(a), float64(b))) }, 3, 4))

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sum(1,2,3,4,5,6))

	Pointer()

	a := 131
	b := 313
	//swap(&a, &b)
	a, b = swap1(a, b)
	fmt.Println(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("unsupported operation:" + op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func div0(a, b int) (q, r int) {
	return a / b, a % b
}

func div1(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

// 这个不知道怎么说
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	funcName := runtime.FuncForPC(p).Name()
	fmt.Println("funcName:" + funcName + "\n")
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for _, number := range numbers {
		s += number
	}
	return s
}

func Pointer()  {

	var a int = 3

	var p *int =&a

	*p = 223
	fmt.Println(a)
}

func swap(a, b *int)  {
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int)  {
	return b, a
}
