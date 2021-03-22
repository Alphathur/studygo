package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	WriteByIoUtil("H:\\go\\main\\2.txt")
	WriteByOs("H:\\go\\main\\3.txt")
}

func WriteByIoUtil(filename string) {
	d := []byte("hello\ngo\n")
	err := ioutil.WriteFile(filename, d, 0644)
	CheckErr(err)
}

func WriteByOs(filename string) {
	f, err := os.Create(filename)
	CheckErr(err)

	defer f.Close()

	//通过f.Write函数写文件
	d2 := []byte("some\n")
	n2, err := f.Write(d2)
	CheckErr(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//通过f.WriteString函数写文件
	n3, err := f.WriteString("writes\n")
	CheckErr(err)
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()

	//通过bufio提供的Writer.WriteString函数写文件
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	CheckErr(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
