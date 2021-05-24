package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan  int
	done func()
}

func channelDemo()  {
	var  wg sync.WaitGroup

	var workers [10]worker

	for i:=0; i<10; i++  {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers  {
		worker.in <- 'a' + i
	}

	for i, worker := range workers  {
		worker.in <- 'A' + i
	}

	wg.Wait()

}

func createWorker(id int, wg *sync.WaitGroup)  worker {
	w := worker{
		in: make(chan  int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func doWorker(id int, w worker)  {
	for n :=range  w.in {
		fmt.Printf("I am workerId:%d,  receive:%c\n", id, n)
		w.done()
	}
}

func main() {
	channelDemo()
}


