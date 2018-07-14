package engine

import (
	"log"

	"github.com/XiaoZhangJian/MeiZiTu/fetcher"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)

		requests = append(requests, parserResult.Requests...)
		if err != nil {
			continue
		}
		for _, item := range parserResult.Items {
			log.Printf("Got item : %v", item)
		}
	}
}

func worker(r Request) (ParserResult, error) {
	// log.Printf("Fetcher %s ", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher error,  r.url : %s , err : %v \t", r.Url, err)

		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
