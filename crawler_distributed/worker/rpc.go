package worker

import "github.com/zoulongbo/learngo/crawler/engine"

type CrawlService struct{}

func (c CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := UnSerializedRequest(req)
	if err != nil {
		return err
	}
	engineResult ,err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializedResult(engineResult)
	return nil
}


