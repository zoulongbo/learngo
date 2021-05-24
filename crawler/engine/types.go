package engine

import "github.com/zoulongbo/learngo/crawler_distributed/config"

type ParserFunc func(content []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Type    string
	Url     string
	Id      string
	Payload interface{}
}


//定义 空解析器
type NilParser struct {
	
}

func (n NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}

//通用解析器方法 接口实现

type FuncParser struct {
	Parser ParserFunc
	Name string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.Parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.Name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser  {
	return &FuncParser{
		Parser: p,
		Name:   name,
	}
}
