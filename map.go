package main

import "fmt"

func main() {
	m := map[string]string{
		"name":"zoulongbo",
		"age":"27",
		"sex":"man",
	}
	fmt.Println("m :" , m)

	m1 := make(map[string]int)		//此时是一个空map map[]

	var m2  map[string]int		//此时是nil 好像类似var定义的 空变量 都是nil

	fmt.Println("m1， m2:", m1, m2)

	for k, _ := range m {		//k在map中是无序的 打出来可能会顺序不一样
		fmt.Printf("m遍历下 k=%s\n", k)
		//fmt.Printf("m遍历下 k=%s v=%s\n", k, v)
	}
	fmt.Println("Getting Value")

	name, ok := m["name"]		//k不存在时 返回一个类型的空值
	fmt.Println("m.name = ", name, ok)


	if m1Name, ok := m1["name"]; ok {
		fmt.Println("m1.name = ", m1Name, ok)

	} else {
		fmt.Println("m1.name不存在")
	}

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("m.name = ", name, ok)

	fmt.Println("m = ", m)


}
