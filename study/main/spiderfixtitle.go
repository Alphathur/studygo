package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"log"
	"time"
)

func main() {
	fixtitle()
}

func fixtitle() {
	db1, err := sql.Open("mysql", "root:mysql520@(127.0.0.1:3306)/demo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db1.Ping(); err != nil {
		log.Fatal(err)
	}

	type Link struct {
		id    int32
		url   string
		title sql.NullString
	}

	rows, err := db1.Query(`SELECT id, url, title FROM sbjnjdaj`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u Link
		err := rows.Scan(&u.id, &u.url, &u.title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u.url)
		title := findtitle(u.url)
		fmt.Println(title)
		db1.Exec("UPDATE sbjnjdaj SET title = ? where id = ?", title, u.id)
		time.Sleep(time.Millisecond * 500)
	}
}

func findtitle(url string) string {
	c := colly.NewCollector()
	var title string
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})
	c.Visit(url)
	return title
}
