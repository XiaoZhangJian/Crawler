package main

import (
	"github.com/XiaoZhangJian/Crawler/engine"
	"github.com/XiaoZhangJian/Crawler/meizitu/parser"
	"github.com/XiaoZhangJian/Crawler/persist"
	"github.com/XiaoZhangJian/Crawler/scheduler"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 10,
		ItemChan:  persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "https://www.dbmeinv.com/",
		ParserFunc: parser.TagList,
	})

	// e.Run(engine.Request{
	// 	Url:        "https://www.dbmeinv.com/dbgroup/show.htm?cid=6",
	// 	ParserFunc: parser.Tag,
	// })

}
