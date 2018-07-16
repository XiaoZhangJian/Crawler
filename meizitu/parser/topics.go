package parser

import (
	"fmt"
	"regexp"

	"github.com/XiaoZhangJian/Crawler/engine"
	"github.com/XiaoZhangJian/Crawler/model"
)

var (
	userNameRe  = regexp.MustCompile(`<a data-author="true" data-name=".*?" href="javascript:;">([^<]+)</a>`)
	avatarRe    = regexp.MustCompile(`<img class=".*?" src="(https://img3.doubanio.com/icon/.*?.jpg)" alt="" />`)
	pushAtRe    = regexp.MustCompile(`<abbr title=".*?"+[^>]*>([^<]+)</abbr>`)
	pushTitleRe = regexp.MustCompile(`<h1 class="media-heading">([^<]+)</h1>`)
	pushImgRe   = regexp.MustCompile(`<img src="([a-zA-z]+://[^\s]*)" width=".*?" alt=".*?" />`)
	pushTextRe  = regexp.MustCompile(`<p>([^<]+)</p>`)

	categoryRe = regexp.MustCompile(`cid=(.*?)">(.*?)</a>`)
)

func Topics(contents []byte, category string, url string) engine.ParserResult {
	topics := model.Topics{}
	topics.UserName = extractString(contents, userNameRe)
	topics.Avatar = extractString(contents, avatarRe)
	topics.PushTitle = extractString(contents, pushTitleRe)
	topics.PushAt = extractString(contents, pushAtRe)
	pushText := pushTextRe.FindAllSubmatch(contents, -1)
	for _, t := range pushText {
		text := string(t[1])
		if text != "我就是测试回复一下的" {
			topics.PushText = append(topics.PushText, text)
		}

	}

	topics.Category = category
	topics.TopicesUrl = url

	pushImg := pushImgRe.FindAllSubmatch(contents, -1)
	for _, m := range pushImg {
		topics.PushImgs = append(topics.PushImgs, string(m[1]))
	}

	categoryRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{
		Items: []interface{}{topics},
	}

	fmt.Printf("============  ====  %+v", topics)

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
