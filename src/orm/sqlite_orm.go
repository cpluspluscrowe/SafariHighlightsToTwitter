package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

}

type Highlight struct {
	Text   string `gorm:"not null;unique"`
	Url    string
	Book   string
	Posted int
}

var prodDbName string = "highlightTweets.db"

func Insert(highlight Highlight) {
	insert(highlight, prodDbName)
}

func GetUnpostedHighlights() []Highlight {
	return getUnposted(prodDbName)
}

func SetAllHighlightsAsPosted() {
	setAllHighlightsAsPosted(prodDbName)
}

func setAllHighlightsAsPosted(dbName string) {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var highlights []Highlight
	db.Model(&highlights).Update("Posted", 1)
}

func insert(highlight Highlight, dbName string) {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(highlight)
}

func getUnposted(dbName string) []Highlight {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var highlights []Highlight
	db.Where("posted = ?", 0).Find(&highlights)
	return highlights
}
