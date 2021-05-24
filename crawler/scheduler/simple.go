package scheduler

import (
	"github.com/zoulongbo/learngo/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}


func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) WorkChan() (c chan engine.Request)  {
	return s.workerChan
}

func (s *SimpleScheduler) WorkReady(w chan engine.Request) {}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

