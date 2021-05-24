package parser

import (
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler/model"
	"io/ioutil"
	"testing"
)

var profile = model.Profile{
	Name:          "昨夜星辰",
	Gender:        "女士",
	Age:           50,
	Height:        171,
	Weight:        0,
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

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "昨夜星辰", "http://album.zhenai.com/u/1496876501")
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element ; but was %v", result.Items)
	}
	actual := result.Items[0]
	if actual != item {
		t.Errorf("excepted %v, but was %v", actual, item)
	}
}
