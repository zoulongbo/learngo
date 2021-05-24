package main

import "fmt"

func printSlice(sli []int)  {
	fmt.Printf("sli=%v, len(sli) = %d, cap(sli) = %d \n", sli, len(sli), cap(sli))
}

func main() {
	fmt.Println("Create Slice")
	var sli []int
	for i := 0; i < 40; i++ {
		sli = append(sli, 2*i +1)
		//printSlice(sli)
	}

	s1 := []int{1,2,3,4}
	printSlice(s1)

	s2 := make([]int, 10)
	printSlice(s2)

	s3 := make([]int, 5, 10)
	printSlice(s3)
	fmt.Println("Copy Slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Delete Slice Element")
	//删除第三个
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	printSlice(s2[2:10])
}
