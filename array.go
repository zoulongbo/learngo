package main

import "fmt"

func main() {
	var arr1 [5]int

	arr2 := [3]int{1, 2, 3}

	arr3 := [...]int{1,3,4,5,1}

	var arr4[3][2]int

	arr5 :=[2][3]int{
		{1, 3, 4},
		{12, 4, 9},
	}

	fmt.Println(arr1, arr2, arr3, arr4, arr5)

	for _,v := range arr3 {
		fmt.Println("我是range来的:", v)
	}

	for i:= 0; i < len(arr2) ; i++ {
		fmt.Println("我是for来的:", arr2[i])
	}

	printArray(arr3)

}

func printArray(arr [5]int)  {
	arr[0] = 100
	for i,v := range arr {
		fmt.Printf("我是range来的: i=%d,v=%d\n", i,v)
	}
}
