package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Movie struct {
	Poster string
}

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		scanner := bufio.NewScanner(in)
		var title string
		for scanner.Scan() {
			title = scanner.Text()
			DownloadImage(title)
		}
	}
}

//练习 4.13： 使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。
func DownloadImage(title string) {
	//url := "http://www.omdbapi.com/?t=" + title + "&apikey=5acab00e"
	//resp, err := http.Get(url)
	//don't use http.Get(url) directly since the query may contain empty string in parameter
	req, err := http.NewRequest("GET", "http://www.omdbapi.com/", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("apikey", "5acab00e")
	q.Add("t", title)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	decoder := json.NewDecoder(resp.Body)
	var movie Movie
	err1 := decoder.Decode(&movie)
	if err1 != nil {
		panic(err1)
	}
	imagLink := movie.Poster
	imgResp, err := http.Get(imagLink)
	defer imgResp.Body.Close()

	filename := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	pwd, _ := os.Getwd()
	//open a file for writing
	file, err := os.Create(pwd + "\\study\\res\\" + filename + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, imgResp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}
