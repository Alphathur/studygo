package main

import (
	"database/sql"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	storeAllImage()
}

func storeAllImage() {
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
		title string
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
		time.Sleep(time.Millisecond * 500)
		var path = "F:\\视频\\sbjnjdaj\\" + u.title
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0777)
		}
		findAllImages(u.url, path)
		fmt.Printf("SAVED ! TITLE : %s , URL : %s \n", u.title, u.url)
	}
}

func findAllImages(url string, folder string) {
	c := colly.NewCollector()
	var index = 1
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		imgsrc := e.Attr("src")
		var filename = folder + "\\" + strconv.Itoa(index) + ".jpg"
		fmt.Println("filename:" + filename)
		saveImage(imgsrc, filename)
		index = index + 1
	})
	c.Visit(url)
}

func saveImage(imgsrc string, filename string) {
	c := colly.NewCollector()
	c.OnResponse(func(response *colly.Response) {
		response.Save(filename)
	})
	c.Visit(imgsrc)
}
