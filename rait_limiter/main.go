package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"time"
)

const (
	LIMITER = 1
	BURST = 5
)

func main() {
	log.Printf("limit %d , burst %d", LIMITER, BURST)
	l := rate.NewLimiter(LIMITER,BURST)

	for i:=1; i < 100 ; i++  {
		log.Println(i, ": before wait")
		// 阻塞等待 令牌
		c, _ := context.WithTimeout(context.Background(), 2 * time.Second)
		if err := l.Wait(c); err != nil {
			log.Println(i, ": limiter wait err:" + err.Error())
		}

		log.Println(i, ": after wait")

		//返回下一个令牌发放时间
		r := l.Reserve()
		log.Printf("%d : limiter delay %v", i, r.Delay())

		//判断当前是否拿令牌
		if !l.Allow() {
			log.Println(i, ": limiter no token")
		}

		time.Sleep(200 * time.Millisecond)
		log.Println(i, ":", time.Now().Format("2006-01-02 15:04:05.000"))
		log.Println()

	}
}
