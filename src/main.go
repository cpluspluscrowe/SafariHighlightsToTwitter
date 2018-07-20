package main

import (
	"fmt"
	"highlightDb"
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
	highlightDb.InsertHighlights(highlights)

	dbHighlights := highlightDb.GetUnpostedHighlights()
	fmt.Printf("Number of highlights to post: %d\n", len(dbHighlights))
	for _, highlight := range dbHighlights {
		twitter.Tweet(highlight.Text)
		highlightDb.SetHighlightAsPosted(highlight.Text)
	}
}
