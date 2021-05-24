package client

import (
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	"github.com/zoulongbo/learngo/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(r engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializedRequest(r)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.UnSerializedParseResult(sResult), nil
	}
}
