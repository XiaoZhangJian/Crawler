package parser

import (
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

func Profile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.UserName = extractString(contents, userNameRe)
	profile.Avatar = extractString(contents, avatarRe)
	profile.PushTitle = name //extractString(contents, pushTitleRe)
	profile.PushAt = extractString(contents, pushAtRe)
	pushText := pushTextRe.FindAllSubmatch(contents, -1)
	for _, t := range pushText {
		text := string(t[1])
		if text != "我就是测试回复一下的" {
			profile.PushText = append(profile.PushText, text)
		}

	}

	pushImg := pushImgRe.FindAllSubmatch(contents, -1)
	for _, m := range pushImg {
		profile.PushImgs = append(profile.PushImgs, string(m[1]))
	}

	categoryRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}

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
