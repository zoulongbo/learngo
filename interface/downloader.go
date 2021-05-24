package main

import (
	"fmt"
	"github.com/zoulongbo/learngo/interface/test"
)

func getRetriever() retriever {
	return test.Retriever{}
}

type retriever interface {
	Get(url string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Printf("%s \n", r.Get("http://www.imooc.com"))
}
