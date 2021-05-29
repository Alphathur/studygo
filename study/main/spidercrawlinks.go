package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"time"
)

var repeatTimes = 0

func main() {
	crawAllLinks()
}

func crawAllLinks() {
	defer func() {
		//恢复程序的控制权
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	c := colly.NewCollector()
	// Find and visit all links
	c.OnHTML(".media", func(e *colly.HTMLElement) {
		var jumps = e.Attr("onclick")
		s := strings.ReplaceAll(jumps, "jump", "")
		s1 := strings.ReplaceAll(s, "('", "")
		s2 := strings.ReplaceAll(s1, "')", "")
		if repeatTimes == 100 {
			log.Println("RepeatTimes is 100, Stop")
		}
		saveToDB(s2)
		time.Sleep(time.Millisecond * 500)
		e.Request.Visit(s2)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("http://sbjnjdaj.zghnhxs.com:666/qq.php?t=162122794123194")
}

func saveToDB(url string) {
	defer func() {
		//恢复程序的控制权
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	db1, err := sql.Open("mysql", "root:mysql520@(127.0.0.1:3306)/demo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db1.Ping(); err != nil {
		log.Fatal(err)
	}
	createdAt := time.Now()
	result, err := db1.Exec(`INSERT INTO sbjnjdaj (url, created_at) VALUES (?, ?)`, url, createdAt)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			repeatTimes += 1
			log.Println("Repeat!")
		}
	}
	if err == nil {
		repeatTimes = 0
	}
	id, err := result.LastInsertId()
	fmt.Println(id)
}
