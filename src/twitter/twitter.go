package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
	"os"
)

func Tweet(tweetText string) {
	httpClient := getClient()

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	//	tweet, resp, err :=
	_, _, err := client.Statuses.Update(tweetText, nil)
	if err != nil {
		panic(err)
	}
}

func FakeTweet(tweetText string) {
	fmt.Println(tweetText)
}

func GetTweets(highlight string) {
	httpClient := getClient()
	client := twitter.NewClient(httpClient)

	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: highlight,
	})
	fmt.Println(search, resp, err)
}

func Search(look4 string) {
	twitterClient := getTwitterClient()
	search, resp, err := twitterClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: "Golang",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(search, resp, err)
}

func getClient() *http.Client {
	consumerKey := os.Getenv("consumerKey")
	consumerSecret := os.Getenv("consumerSecret")
	accessToken := os.Getenv("accessToken")
	accessSecret := os.Getenv("accessSecret")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return httpClient
}

func getTwitterClient() *twitter.Client {
	httpClient := getClient()
	client := twitter.NewClient(httpClient)
	return client
}
