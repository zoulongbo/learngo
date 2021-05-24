package main

import (
	"fmt"
	"github.com/zoulongbo/learngo/queue"
)

func main() {
	queue := queue.List{1}
	fmt.Print(queue.IsEmpty())
	queue.Push(2, 3, 4)
	fmt.Println(queue)
	fmt.Println(queue.Pop())
	fmt.Println(queue)

	queue.Push("aa","bb", "cc")
	for {
		fmt.Println(queue.Pop())
		if queue.IsEmpty() {
			break
		}
	}
}
