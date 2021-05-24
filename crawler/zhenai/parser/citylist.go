package parser

import (
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		//result.Items = append(result.Items, "City :"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]) + "/nv",
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
