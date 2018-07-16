package persist

import (
	"context"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

// 存储到数据库中
func ItemSaver() (chan interface{}, error) {

	out := make(chan interface{})
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item #%d : %+v", itemCount, item)
			itemCount++
			_, err := save(client, item)
			if err != nil {
				log.Printf("Item Saver err %v : %v", item, err)
				continue
			}

		}
	}()

	return out, nil
}

func save(client *elastic.Client, item interface{}) (id string, err error) {
	resp, err := client.Index().Index("doubanmeizi").Type("meizi").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
