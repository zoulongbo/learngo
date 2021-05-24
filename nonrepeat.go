package main

import "fmt"

func main() {
	fmt.Println(lengthOfNonRepeatSubstr("asdasdweffqefqw"))
	fmt.Println(lengthOfNonRepeatSubstr("asd	we3e123r13ra"))
	fmt.Println(lengthOfNonRepeatSubstr("1234123f1ef34r"))
	fmt.Println(lengthOfNonRepeatSubstr("231234213f3e"))
	fmt.Println(lengthOfNonRepeatSubstr("1ecf423rcf42ef1erc1"))
}

func lengthOfNonRepeatSubstr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		//fmt.Printf("i:%d ch:%b  start:%d  maxLenngth:%d \n", i, ch, start, maxLength)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}
	return maxLength
}