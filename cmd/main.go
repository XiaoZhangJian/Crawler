package main

import (
	"github.com/XiaoZhangJian/Crawler/engine"
	"github.com/XiaoZhangJian/Crawler/meizitu/parser"
	"github.com/XiaoZhangJian/Crawler/persist"
	"github.com/XiaoZhangJian/Crawler/scheduler"
)

func main() {

	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 10,
		ItemChan:  itemChan,
	}
	e.Run(engine.Request{
		Url:        "https://www.dbmeinv.com/",
		ParserFunc: parser.TagList,
	})

}
