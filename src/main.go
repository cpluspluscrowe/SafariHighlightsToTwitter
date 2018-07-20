package main

import (
	"fmt"
	"orm"
	"safari"
	"strings"
	"time"
	"twitter"
)

func doEvery(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}

func main() {
	doEvery(12000*time.Millisecond, postHighlightsToTwitter)
}

func postHighlightsToTwitter() {
	highlights := safari.GetSafariHighlights()
	for _, highlight := range highlights {
		fmt.Println(highlight)
		orm.Insert(highlight)
	}

	//	orm.SetAllHighlightsAsPosted()
	unpostedHighlights := orm.GetUnpostedHighlights()
	fmt.Printf("Number of highlights to post: %d\n", len(unpostedHighlights))
	for _, highlight := range unpostedHighlights {
		twitter.FakeTweet(highlight.Text)
		orm.SetAsPosted(highlight)
	}
}

func formatHighlightForPost(safariHighlights []orm.Highlight) []string {
	highlights := []string{}
	for _, highlight := range safariHighlights {
		citation := strings.TrimSpace(strings.Replace(highlight.Url, "\n", "", 1))
		highlights = append(highlights, highlight.Text+"\n- \""+citation+"\"")
	}
	return highlights
}
