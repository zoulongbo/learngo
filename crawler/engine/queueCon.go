package engine

import (
	"log"
)


func (e *ConcurrentEngine) QueueRun(seeds ...Request) {
	//输出队列
	out := make(chan ParseResult)
	//配置初始化请求队列给 调度器
	e.Scheduler.Run()
	//分配执行线程/协程
	for i := 0; i < e.WorkCount; i++ {
		e.QCreateWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	//提交请求给调度器
	for _, r := range seeds {
		////这里url去重
		//if isDuplicate(r.Url) {
		//	log.Printf("Duplicate Request Url : %s", r.Url)
		//}
		e.Scheduler.Submit(r)
	}
	for {
		//输出 获取到的item
		result := <-out
		for _, item := range result.Items {
			go func() {e.ItemChan <- item}()
		}
		//再将 获取到的item.request 输入
		for _, request := range result.Requests {
			////这里url去重
			//if isDuplicate(request.Url) {
			//	log.Printf("Duplicate Request Url : %s", request.Url)
			//	continue
			//}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) QCreateWorker(in chan Request, out chan ParseResult, ready ReadyNotify) {
	go func() {
		for {
			ready.WorkReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				log.Printf("worker result err:%v", err)
				continue
			}
			out <- result
		}
	}()
}

////hash map
//var visitedUrls = make(map[string]bool)
//
//func isDuplicate(url string) bool {
//	if visitedUrls[url] {
//		return true
//	}
//	visitedUrls[url] = true
//	return false
//}
