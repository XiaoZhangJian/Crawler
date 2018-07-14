package persist

import (
	"context"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item #%d : %+v", itemCount, item)
			itemCount++

			// _, err := save(item)
			// if err != nil {
			// 	log.Printf("Item Saver err %v : %v", item, err)
			// 	continue
			// }

		}
	}()

	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().Index("doubanmeizi").Type("meizi").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	// fmt.Printf("%+v", resp)
	return resp.Id, nil
}
