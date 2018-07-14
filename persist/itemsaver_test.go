package persist

import (
	"context"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"

	"github.com/XiaoZhangJian/MeiZiTu/model"
)

func TestSaver(t *testing.T) {

	mode := model.Profile{
		UserName:  "不存在的小马甲",
		Avatar:    "https://img3.doubanio.com/icon/up63849034-4.jpg",
		PushTitle: "点开便知",
		PushAt:    "2018-05-09 15:01:45.0",
		PushText:  []string{"健身大概是为了吃更多的好吃的哈哈哈哈哈哈哈嘻嘻嘻"},
		PushImgs:  []string{"https://wx2.sinaimg.cn/large/0060lm7Tgy1fratapfh96j30dw072dgn.jpg"},
	}

	id, err := save(mode)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("doubanmeizi").Type("meizi").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", resp)
	var tiezi model.Profile
	err = json.Unmarshal(*resp.Source, &tiezi)
	if err != nil {
		panic(err)
	}

}
