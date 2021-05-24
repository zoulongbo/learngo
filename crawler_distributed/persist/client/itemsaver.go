package client

import (
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	"github.com/zoulongbo/learngo/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver : Got item  $%d, %v", itemCount, item)
			itemCount++
			//Call Rpc
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil || result != "OK" {
				log.Printf("Rpc Item Saver:error saving item %v, %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}