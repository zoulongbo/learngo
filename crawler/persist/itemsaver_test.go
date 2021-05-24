package persist

import (
	"github.com/olivere/elastic/v7"
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler/model"
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

func TestSave(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		t.Error(err)
	}
	err = Save(client, item, "dating_profile")
	if err != nil {
		t.Error(err)
	}
	expected, err := getById(client, item.Id, item.Type)
	if err != nil {
		t.Error(err)
	}

	actualProfile, _ := model.FromJsonObj(expected.Payload)

	expected.Payload = actualProfile

	if expected != item {
		t.Errorf("Got %v, expected %v", expected, item)
	}
}
