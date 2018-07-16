package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WokerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

/*
	并发版实现
*/
func (c *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParserResult)
	c.Scheduler.Run()
	for i := 0; i < c.WorkCount; i++ {
		createWorker(c.Scheduler.WokerChan(), out, c.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicte(r.Url) {
			log.Printf("Duplicte request %s", r.Url)
			continue
		}
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {

			go func() {
				c.ItemChan <- item
			}()
		}
		// URL 去重
		for _, request := range result.Requests {
			if isDuplicte(request.Url) {
				log.Printf("Duplicte request %s", request.Url)
				continue
			}
			c.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrl = make(map[string]bool)

func isDuplicte(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}
