package main

import (
	"bufio"
	"fmt"
	flb2 "github.com/zoulongbo/learngo/functional/flb"
	"os"
)

func tryDefer()  {
	defer fmt.Println(1)		//栈		先进后出
	defer fmt.Println(2)
	 fmt.Println(3)
}

func writeFile(filename string) {
	/*file, err := os.Create(filename)
	if err != nil  {
		panic(err)
	}*/
	file, err :=  os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("Error: %s , %s, %s", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	flb := flb2.Fibonacci()
	for i:=0; i<20; i++ {
		fmt.Fprintln(writer, flb())
	}

}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
