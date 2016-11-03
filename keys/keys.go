package keys

import (
	"github.com/ChimeraCoder/anaconda"
	. "os"
)

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(Getenv("TWITTER_CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(Getenv("TWITTER_ACCESS_TOKEN"), Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
}
