package safari

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

type Highlight struct {
	Text   string
	Url    string
	Book   string
	Posted int
}

func GetSafariHighlights() []string {
	highlights := getSafariHighlights()
	withCitations := addCitationToHighlights(highlights)
	return withCitations
}

func getSafariHighlights() []Highlight {
	url := "https://www.safaribooksonline.com/u/8aaaf02c-85f9-42f8-8628-3e0c83d24fd6/"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	highlights, err := extractHighlights(resp.Body)
	if err != nil {
		panic(err)
	}
	return highlights
}

func addCitationToHighlights(safariHighlights []Highlight) []string {
	highlights := []string{}
	for _, highlight := range safariHighlights {
		citation := strings.TrimSpace(strings.Replace(highlight.Url, "\n", "", 1))
		highlights = append(highlights, highlight.Text+"\n- \""+citation+"\"")
	}
	return highlights
}

func extractHighlights(body io.Reader) ([]Highlight, error) {
	highlights := []Highlight{}
	toPrint := false
	z := html.NewTokenizer(body)
	toGetSource := false
	highlightText := ""
	link := ""
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return highlights, nil
		case html.TextToken:
			if toGetSource {
				//source := string(z.Text())
				highlights = append(highlights, Highlight{Text: highlightText, Url: link})
				toGetSource = false
			}
			if toPrint {
				highlightText = string(z.Text())
				toPrint = false
			}
		case html.StartTagToken, html.EndTagToken:
			_, value, _ := z.TagAttr()
			if string(value) == "t-annotation-quote annotation-quote" {
				toPrint = true
			}
			if string(value) == "t-annotation-archive-title" {
				t := z.Token()
				for _, a := range t.Attr {
					if a.Key == "href" {
						link = "https://www.safaribooksonline.com" + a.Val
					}
				}
				toGetSource = true
			}
		}
	}
}

//
