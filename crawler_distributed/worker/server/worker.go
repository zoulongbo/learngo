package main

import (
	"flag"
	"fmt"
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	"github.com/zoulongbo/learngo/crawler_distributed/rpcsupport"
	"github.com/zoulongbo/learngo/crawler_distributed/worker"
	"log"
)

var port  = flag.Int("port", config.WorkPort0, "the port for me to listen on, there is an default one if you don`t input")

func main() {
	flag.Parse()

	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
