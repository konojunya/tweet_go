package main

import (
	. "./keys"
	. "fmt"
)

func main() {

	api := GetTwitterApi()

	text := "Hello world"
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}

	Print(tweet.Text)
}
