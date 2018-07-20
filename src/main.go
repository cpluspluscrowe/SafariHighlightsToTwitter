package main

import (
	"fmt"
	"orm"
	"safari"
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

	unpostedHighlights := orm.GetUnpostedHighlights()
	fmt.Printf("Number of highlights to post: %d\n", len(unpostedHighlights))
	for _, highlight := range unpostedHighlights {
		twitter.Tweet(highlight.Text)
	}
}
