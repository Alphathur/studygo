package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := "H:\\go\\main\\1.txt"
	ReadByIoUntil(file)
	ReadByOs(file)
}

func ReadByIoUntil(filename string) {
	fmt.Println("-------通过ioutil.ReadFile读取---------")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
	fmt.Println(string(data))
}

func ReadByOs(filename string) {
	fmt.Println("-------通过os.Open读取---------")
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}

	//采用bufio读取文件内容
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())
	}

	//采用ioutil读取文件内容，由于流是一次性的，所以上面的代码输出后，下面的Println将不再打印任何东西
	contents1, err := ioutil.ReadAll(f)
	fmt.Println(contents1)

	f.Close()
}
