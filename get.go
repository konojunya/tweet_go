package main

import (
	. "./keys"
	. "fmt"
	"net/url"
)

func main() {

	api := GetTwitterApi()

	v := url.Values{}
	v.Set("count","100")

	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		Println("tweet: ", tweet.Text)
	}

}
