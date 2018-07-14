package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/XiaoZhangJian/Crawler/model"
)

func TestProfile(t *testing.T) {
	content, err := ioutil.ReadFile("./profile_data_test.html")

	if err != nil {
		panic(err)
	}

	result := Profile(content)
	if len(result.Items) != 1 {
		log.Printf("err: %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	fmt.Println(profile.UserName)
	fmt.Println(profile.Avatar)
	fmt.Println(profile.PushTitle)
	fmt.Println(profile.PushAt)
	fmt.Println(string(profile.PushImgs))
	fmt.Println(profile.PushText)

}
