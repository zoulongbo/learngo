package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkCount        int
	ItemChan         chan Item
	CacheDriver      CacheDriver
	RequestProcessor Processor
}

type Processor func(r Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotify
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotify interface {
	WorkReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//输出队列
	out := make(chan ParseResult)
	//配置初始化请求队列给 调度器
	e.Scheduler.Run()
	//分配执行线程/协程
	for i := 0; i < e.WorkCount; i++ {
		e.createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	//提交请求给调度器
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		//输出 获取到的item
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		//再将 获取到的item.request 输入
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotify) {
	go func() {
		for {
			ready.WorkReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
