package main

import "fmt"

func main() {
	fmt.Println(lengthOfNonRepeatSubstr("1ecf423rcf42ef1erc1"))
	fmt.Println(lengthOfNonRepeatSubstr("我的山鸡舞镜我饿收到欧水情人我去玩儿欧水啊"))
	fmt.Println(lengthOfNonRepeatSubstr("11121312341234132413413424213413"))
	fmt.Println(lengthOfNonRepeatSubstr("1122231"))
	fmt.Println(lengthOfNonRepeatSubstr(" "))
}

func lengthOfNonRepeatSubstr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
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