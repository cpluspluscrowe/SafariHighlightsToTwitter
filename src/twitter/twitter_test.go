package twitter

import (
	"fmt"
	"os"
	"testing"
)

func TestGetTweets(t *testing.T) {
	GetTweets("other")

}

func TestSearch(t *testing.T) {
	Search("golang")
}

func TestEnv(t *testing.T) {
	consumerKey := os.Getenv("consumerKey")
	consumerSecret := os.Getenv("consumerSecret")
	accessToken := os.Getenv("accessToken")
	accessSecret := os.Getenv("accessSecret")
	fmt.Println(consumerKey)
	fmt.Println(consumerSecret)
	fmt.Println(accessToken)
	fmt.Println(accessSecret)
}
