package worker

import (
	"errors"
	"fmt"
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler/zhenai/parser"
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializedRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializedResult(res engine.ParseResult) ParseResult {
	result := ParseResult{
		Items:res.Items,
	}
	for _ , request :=  range res.Requests {
		result.Requests = append(result.Requests, SerializedRequest(request))
	}
	return result
}

func UnSerializedRequest(request Request) (engine.Request, error) {
	parser, err := unSerializedParser(request.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    request.Url,
		Parser: parser,
	}, nil
}

func unSerializedParser(p SerializedParser) (engine.Parser, error)  {
	switch p.Name {
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseProfile:
		if username, ok := p.Args.(string); ok {
			return parser.NewProfileParser(username), nil
		} else {
			return nil, fmt.Errorf("invalid arg :%v", p.Args )
		}
	case config.ProfileMockParser:
		if username, ok := p.Args.(string); ok {
			return parser.NewProfileMockParser(username), nil
		} else {
			return nil, fmt.Errorf("invalid arg :%v", p.Args )
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser")
	}
}

func UnSerializedParseResult(res ParseResult) engine.ParseResult {
	result :=  engine.ParseResult{
		Items:    res.Items,
	}
	for _, req := range res.Requests {
		request, err := UnSerializedRequest(req)
		if err != nil {
			log.Printf("error unSerialized request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}
