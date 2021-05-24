package persist

import (
	"github.com/olivere/elastic/v7"
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	if err == nil {
		*result = "OK"
	}
	return err
}
