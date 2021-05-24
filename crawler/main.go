package main

import (
	"flag"
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler/scheduler"
	"github.com/zoulongbo/learngo/crawler/zhenai/parser"
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	itemSaver "github.com/zoulongbo/learngo/crawler_distributed/persist/client"
	"github.com/zoulongbo/learngo/crawler_distributed/rpcsupport"
	worker "github.com/zoulongbo/learngo/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var(
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String("worker_hosts", "", "worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	//engine.SimpleEngine{}.Run(engine.Request{
	//e :=engine.Concurrent{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkCount: 10,
	//}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	//itemChan, err := persist.ItemSaver("dating_profile")
	itemChan, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	workerPool := createWorkClientPool(strings.Split(*workerHosts, ","))
	worker	:= worker.CreateProcessor(workerPool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkCount:        10,
		ItemChan:         itemChan,
		CacheDriver:      engine.CreateCacheDriver(),
		RequestProcessor: worker,
	}

	e.QueueRun(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
	//e.QueueRun(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/xiamen",
	//	ParserFunc: parser.ParseCity,
	//})
}

func createWorkClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcsupport.NewClient(host)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connecting to %s", host)
		} else {
			log.Printf("error connecting to %s : %v", host, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
