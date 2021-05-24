package main

import (
	"fmt"
	"math"
)

var outSide = 3333

func main() {
	fmt.Println("hello world")
	fmt.Println(outSide)
	variable()
	variableInit()
	variableTypeDeduction()
	variableShorter()

	consts()

	enums()

	triangle()
}

func variable() {
	var a int
	var b string
	fmt.Printf("a = %d  b=%q \n", a, b)
}

//初始化
func variableInit() {
	//var a, b = 1, 2
	var a, b int = 1, 2
	var c, d string = "c", "d"
	//var c, d  = "c", "d"
	fmt.Println(a, b, c, d)
}

//推断类型
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, "ok", true
	fmt.Println(a, b, c, d)
}

//定义并赋值
func variableShorter() {
	a, b, c, d := 1, 2, "ok", true
	fmt.Println(a, b, c, d)
}

const fileName string = "one.png"

//常量
func consts() {
	//const fileName = "one.png"

	var c int
	//const a, b int = 1, 3
	//c = int(math.Sqrt(float64(a*a + b*b)))
	const a, b = 1, 3

	c = int(math.Sqrt((a*a + b*b)))

	fmt.Println(fileName, c)
}

//枚举
func enums() {
	//普通枚举
	const (
		aa = 0
		bb = 1
		cc = 2
		dd = 3
	)

	//自增枚举
	const (
		cpp = iota //iota=0  其余顺序自增
		_
		java
		php
		golang
	)

	const (
		b = 1 << (10 * iota)
		c
		d
		e
	)

	fmt.Println(aa, bb, cc, dd)
	fmt.Println(cpp, golang, java, php)
	fmt.Println(b, c, d, e)
}


func triangle(){
	var a,b int = 3,4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int  {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}