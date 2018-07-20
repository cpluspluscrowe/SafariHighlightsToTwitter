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
		orm.Insert(highlight)
	}

	//	orm.SetAllHighlightsAsPosted()
	unpostedHighlights := orm.GetUnpostedHighlights()
	fmt.Printf("Number of highlights to post: %d\n", len(unpostedHighlights))
	for _, highlight := range unpostedHighlights {
		formattedPost := formatHighlightForTwitter(highlight)
		twitter.Tweet(formattedPost)
		orm.SetAsPosted(highlight)
	}
}

func formatHighlightForTwitter(highlight orm.Highlight) string {
	formatted := strings.TrimSpace(strings.Replace(highlight.Url, "\n", "", 1))
	return formatted
}
