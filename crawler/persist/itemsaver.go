package persist

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/zoulongbo/learngo/crawler/engine"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver : Got item  $%d, %v", itemCount, item)
			itemCount++
			err := Save(client, item, index)
			if err != nil {
				log.Printf("Item Saver:error saving item %v, %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) (err error) {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.
		Do(context.Background())

	return err
}

func getById(client *elastic.Client,id string, typ string) (item engine.Item, err error) {
	var actual engine.Item
	resp, err := client.Get().
		Index("dating_profile").
		Type(typ).
		Id(id).
		Do(context.Background())
	if err != nil {
		return actual, err
	}
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		return actual, err
	}
	return actual, nil
}
