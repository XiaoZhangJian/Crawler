package parser

import (
	"fmt"
	"regexp"

	"github.com/XiaoZhangJian/Crawler/engine"
)

const (
	TagListRe = `cid=(.*?)">(.*?)</a>`
)

func TagList(contents []byte) engine.ParserResult {
	rep := regexp.MustCompile(TagListRe)
	req := rep.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range req {
		category := string(m[2])
		fmt.Printf("分组：%s \n", category)
		fmt.Printf("URL %s ", fmt.Sprintf("https://www.dbmeinv.com:443/dbgroup/show.htm?cid=%s", m[1]))

		// result.Items = append(result.Items, "Tag --> "+category)
		result.Requests = append(result.Requests, engine.Request{
			Url:        fmt.Sprintf("https://www.dbmeinv.com:443/dbgroup/show.htm?cid=%s", m[1]),
			ParserFunc: Tag,
		})
	}
	return result
}
