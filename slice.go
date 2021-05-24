package main

import "fmt"

func main() {
	arr := [...]int{61, 62, 63, 64, 65, 66, 67}
	arr1 := arr[1:5]
	arr2 := arr[:5]
	s1 := arr[1:]
	s2 := arr[:]

	fmt.Println("我是arr[1:5] :", arr1)
	fmt.Println("我是arr[:5] :", arr2)



	fmt.Println("我是s1 :", s1)
	fmt.Println("我是s2 :", s2)
	fmt.Println("updateSlice(s1):")
	updateSlice(s1)
	fmt.Println("我是s1 :", s1)
	fmt.Println("我是arr :", arr)

	fmt.Println("updateSlice(s2):")
	updateSlice(s2)
	fmt.Println("我是s2 :", s2)
	fmt.Println("我是arr :", arr)

	fmt.Println("ReSlice(s2):")
	s2 = s2[:5]
	fmt.Println("我是s2 = s2[:5] :", s2)
	s2 = s2[2:]
	fmt.Println("我是s2 = s2[2:] :", s2)
	fmt.Println("我是arr :", arr)


	fmt.Println("Slice Extend")
	arr[0] = 200
	fmt.Println("我是arr本初 :", arr)
	ss1 := arr[1:5]
	fmt.Printf("ss1=%v, len(ss1)= %d, cap(ss1) = %d \n", ss1, len(ss1), cap(ss1))
	ss2 := ss1[3:5]
	fmt.Printf("ss2=%v, len(ss2)= %d, cap(ss2) = %d \n", ss2, len(ss2), cap(ss2))
	ss3 := ss2[1:3]		//可向后扩展 向前不行
	fmt.Printf("ss3=%v, len(ss3)= %d, cap(ss3) = %d \n", ss3, len(ss3), cap(ss3))

	fmt.Println("Slice Extend: Add Column")
	ss4 := append(ss3, 10001)
	fmt.Println("append(ss3, 10001) = ss4 = ", ss4)
	ss5 := append(ss4, 20001)
	fmt.Println("append(ss4, 20001) = ss5 = ", ss5)
	ss6 := append(ss5, 30001)
	fmt.Println("append(ss5, 30001) = ss6 = ", ss6)
	fmt.Printf("ss6=%v, len(ss6)= %d, cap(ss6) = %d \n", ss6, len(ss6), cap(ss6))	//copy数组 给更大内存的
	ss61 := ss6[2:8]
	fmt.Println("ss6[2:5] ", ss61)

	fmt.Println("看看ss3 =  ", ss3)
	fmt.Println("看看arr =  ", arr)




}

func updateSlice(s []int)  {
	s[0] = 100
}
