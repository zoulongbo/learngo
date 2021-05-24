package main

import "fmt"

func main() {
	s := "这是中文字符"
	for _, b := range []byte(s) {
		fmt.Printf("%X  ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("  (%d  %X)  ", i, ch)
	}

}
