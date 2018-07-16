package parser

import (
	"fmt"
	"regexp"

	"github.com/XiaoZhangJian/Crawler/engine"
)

var (
	TopicsRe = regexp.MustCompile(`<a href="(https://www.dbmeinv.com:443/dbgroup/[1-9]\d*)" +[^>]*>([^<]+)</a>`)

	NextTagRe = regexp.MustCompile(`<li class="next next_page"><a href="(.*?)" title="(.*?)">[^<]+</a></li>`)
)

func Tag(contents []byte, category string) engine.ParserResult {

	req := TopicsRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range req {

		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(contents []byte) engine.ParserResult {
				return Topics(contents, category, url)
			},
		})
	}

	// 下一页
	next := NextTagRe.FindAllSubmatch(contents, -1)
	for _, n := range next {
		if string(n[2]) == "下一页" {
			url := fmt.Sprintf("https://www.dbmeinv.com:443%s", n[1])
			result.Requests = append(result.Requests, engine.Request{
				Url: url,
				ParserFunc: func(c []byte) engine.ParserResult {
					return Tag(c, category)
				},
			})
		}
	}

	return result
}
