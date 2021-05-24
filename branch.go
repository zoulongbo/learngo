package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	const fileName = "iknow.txt"

	if file, err := ioutil.ReadFile(fileName); err != nil {
		log.Println(err.Error())

	} else {
		fmt.Printf("%s \n", file)
	}

	fmt.Println(switchK(80), switchK(70), switchK(60))
	fmt.Println(switchKeys(0), switchKeys(90), switchKeys(60), switchKeys(101))

}

//条件语句
func switchKeys(score int) string {
	p := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score %d \n", score))
	case score < 60:
		p = "F"
	case score < 80:
		p = "C"
	case score < 90:
		p = "B"
	case score < 100:
		p = "A"
	}
	return p
}

//等式判断
func switchK(score int) string {
	p := ""
	switch score {
	case 60:
		p = "11"
	case 70:
		p = "22"
	case 80:
		p = "33"
	default:
		p = "232"
	}
	return p
}
