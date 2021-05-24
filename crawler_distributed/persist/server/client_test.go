package main

import (
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler/model"
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	"github.com/zoulongbo/learngo/crawler_distributed/rpcsupport"
	"testing"
	"time"
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

func TestItemSaver(t *testing.T) {
	const host  = ":1234"
	//start ItemSaverServer
	//start ItemSaverClient
	//Call Save
	go serveRpc(host, "dating_profile")

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Error(err)
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "OK" {
		t.Errorf("result: %v", result)
	}
	t.Log("result: " + result)
}
