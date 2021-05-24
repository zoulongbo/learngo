package engine

import (
	"github.com/zoulongbo/learngo/crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error)  {
	log.Printf("Fetching %s", r.Url)
	//去重
	//if cacheDriver.IsDuplicate(r.Url) {
	//	return ParseResult{}, errors.New("Duplicate Request Url :" + r.Url)
	//}
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
		log.Printf("Fetcher:error fetching url: %s : %v", r.Url, err)
	}
	return r.Parser.Parse(body, r.Url), nil
}
