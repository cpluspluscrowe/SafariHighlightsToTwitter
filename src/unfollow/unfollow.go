package unfollow

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
	"os"
)

func GetFollowers() {
	client := getTwitterClient()
	followers, _, err := client.Friends.List(&twitter.FriendListParams{})
	users := followers.Users
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user.ID)
	}
}

func ExamineFriendships() {
	client := getTwitterClient()
	friends, _, err := client.Friends.List(&twitter.FriendListParams{})
	users := friends.Users
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		relationship, _, err := client.Friendships.Show(&twitter.FriendshipShowParams{TargetID: user.ID})
		if err != nil {
			panic(err)
		}
		if !relationship.Target.Following {
			fmt.Println("Unfollowing: ", relationship.Target)
			client.Friendships.Destroy(&twitter.FriendshipDestroyParams{UserID: user.ID})
			break
		}
	}

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
