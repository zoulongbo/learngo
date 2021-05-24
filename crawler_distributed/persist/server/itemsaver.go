package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	"github.com/zoulongbo/learngo/crawler_distributed/persist"
	"github.com/zoulongbo/learngo/crawler_distributed/rpcsupport"
	"log"
)

var port  = flag.Int("port", config.ItemSaverPort, "the port for me to listen on, there is an default one if you don`t input")

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}


func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{Client:client, Index:index})
}
