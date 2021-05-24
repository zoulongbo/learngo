package main

import (
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	"github.com/zoulongbo/learngo/crawler_distributed/rpcsupport"
	"github.com/zoulongbo/learngo/crawler_distributed/worker"
	"log"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host  = ":8800"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url:    "http://album.zhenai.com/u/1496876501",
		Parser: worker.SerializedParser{
			Name: "ProfileMockParser",
			Args: "昨夜星辰",
		},
	}

	var result worker.ParseResult

	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		log.Println(result)
	}
}
