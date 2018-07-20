package highlightDb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Highlight struct {
	Text string
	Id   []uint8
}

func InsertHighlights(highlights []string) {
	db := getDatabaseDriver()
	defer db.Close()
	createHighlightTable(db)
	for _, highlight := range highlights {
		insertHighlight(db, highlight)
	}
}

func getDatabaseDriver() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	checkErr(err)
	return db
}

func GetUnpostedHighlights() []Highlight {
	db, err := sql.Open("sqlite3", "./database.db")
	defer db.Close()
	createHighlightTable(db)
	rows, err := db.Query("SELECT uid,text FROM new_posts")
	defer rows.Close()

	highlights := []Highlight{}
	highlight := Highlight{}
	for rows.Next() {
		err = rows.Scan(&highlight.Id, &highlight.Text)
		checkErr(err)

		highlights = append(highlights, highlight)
	}
	return highlights
}

func GetPostedHighlights() []Highlight {
	db, err := sql.Open("sqlite3", "./database.db")
	defer db.Close()
	createHighlightTable(db)
	rows, err := db.Query("SELECT uid,text FROM posted")
	defer rows.Close()

	highlights := []Highlight{}
	highlight := Highlight{}
	for rows.Next() {
		err = rows.Scan(&highlight.Id, &highlight.Text)
		checkErr(err)

		highlights = append(highlights, highlight)
	}
	return highlights
}

func SetHighlightAsPosted(text string) {
	db := getDatabaseDriver()
	statement, _ := db.Prepare("insert into posted(text) values(?)")
	statement.Exec(text)

	statement, _ = db.Prepare("delete from new_posts where text = ?")
	statement.Exec(text)
}

func insertHighlight(db *sql.DB, highlightText string) {
	statement, _ := db.Prepare("INSERT INTO new_posts(text) select ? where not exists(select * from posted where text = ?);")
	statement.Exec(highlightText, highlightText)
}

func RemoveHighlightFromDatabase(highlightText string) {
	db := getDatabaseDriver()
	statement, _ := db.Prepare("delete from posted where text = ?")
	statement.Exec(highlightText)
}

func createHighlightTable(db *sql.DB) {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS 'new_posts' (
		        	'uid' INTEGER PRIMARY KEY AUTOINCREMENT,
			        'text' TEXT UNIQUE NOT NULL
				);`)
	_, err = stmt.Exec()
	checkErr(err)
	stmt, err = db.Prepare(`CREATE TABLE IF NOT EXISTS 'posted' (
		        	'uid' INTEGER PRIMARY KEY AUTOINCREMENT,
			        'text' TEXT UNIQUE NOT NULL
				);`)
	_, err = stmt.Exec()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
