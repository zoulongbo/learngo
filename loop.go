package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(convertToBin(7), convertToBin(13), convertToBin(41231231), convertToBin(0))

	//类似php foreach
	forRange()

	printFile("iknow.txt")

	// 类型while
	//forever()
}

func convertToBin(n int) string {
	p := ""
	if n == 0 {
		return strconv.Itoa(n)
	}
	for ; n > 0; n /= 2 {
		p = strconv.Itoa(n%2) + p
	}
	return p
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	//省略起始条件 省略 递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("1+1=1+1")
	}
}

func forRange() {
	arr := make(map[string]string)

	arr["a"] = "aaa"
	arr["b"] = "bbb"
	arr["c"] = "ccc"

	for _, v := range arr {
		fmt.Println(v)
	}
}
