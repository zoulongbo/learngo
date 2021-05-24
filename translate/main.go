package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type ParseResult []string

var profileRe  = regexp.MustCompile(`=> "([^"]+)"`)
var validateRe  = regexp.MustCompile(`=> ":attribute ([^"]+)",`)

var readPath = "$PWD/translate/"
var writePath = "$PWD/translate/"
var readFiles = []string{"errors-common.php","fields-common.php","errors-mall.php","fields-mall.php"}


func Parse(fileName string) ParseResult {
	file, err := ioutil.ReadFile(readPath + fileName)
	if err != nil {
		panic(err)
	}
	matches := profileRe.FindAllSubmatch(file, -1)
	result := ParseResult{}
	for _, m := range matches {
		result = append(result, string(m[1]))
	}
	return result
}

func main() {
	for _, fileName := range readFiles {
		result := Parse(fileName)
		WriteMapToFile(result, writePath + fileName)
	}
	//validate := Parse("validation.php")
	//WriteMapToFile(validate, writePath + "validation.php")

	fmt.Println("ok")
}

func WriteMapToFile(m []string, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, v := range m {
		fmt.Fprintln(w, v)
	}
	return w.Flush()
}