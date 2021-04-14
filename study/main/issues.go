package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"time"
)

//练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
func main() {
	lessYear := LessYear()
	fmt.Println(lessYear)

	lessMonth := LessMonth()
	fmt.Println(lessMonth)

	result, err := github.SearchIssues([]string{"kubelet"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	var resMap = make(map[string][]*github.Issue)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)

		if item.CreatedAt.After(lessMonth) { //within a month
			resMap["WITHMONT"] = append(resMap["WITHMONT"], item)
		} else if item.CreatedAt.Before(lessMonth) && item.CreatedAt.After(lessYear) { //within a year and beyond a month
			resMap["WITHYEAR"] = append(resMap["WITHYEAR"], item)
		} else {
			resMap["BEYONDYEAR"] = append(resMap["BEYONDYEAR"], item)
		}
	}
	for s, issues := range resMap {
		fmt.Printf("%s 's issues count is %d \n", s, len(issues))
	}
}

func LessMonth() time.Time {
	now := time.Now()
	return now.AddDate(0, -1, 0)
}

func LessYear() time.Time {
	now := time.Now()
	return now.AddDate(-1, 0, 0)
}
