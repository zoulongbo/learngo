package view

import (
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler/frontend/model"
	common "github.com/zoulongbo/learngo/crawler/model"
	"os"
	"testing"
)

var profile = common.Profile{
	Name:          "昨夜星辰",
	Gender:        "女士",
	Age:           50,
	Height:        171,
	Weight:        50,
	Income:        "3001-5000元",
	Marriage:      "离异",
	Education:     "高中及以下",
	Occupation:    "销售总监",
	Native:        "北京",
	Constellation: "魔羯座",
	House:         "已购房",
	Car:           "已买车",
}

var item = engine.Item{
	Type:    "zhenai",
	Url:     "http://album.zhenai.com/u/1496876501",
	Id:      "1496876501",
	Payload: profile,
}

var filename  = "template.html"
var copyFilename  = "template_test.html"

func TestSearchResultView_Render(test *testing.T) {
	view := CreateSearchResultView(filename)
	pages := model.SearchResult{}
	pages.Hits = 12
	pages.Start = 1

	for i:=0; i < 10; i ++  {
		pages.Items = append(pages.Items, item)
	}

	out, err := os.Create(copyFilename)
	err = view.Render(out, pages)
	if err != nil {
		panic(err)
	}
}
