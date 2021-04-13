package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	_ "unicode/utf8"
)

func main() {
	ages := make(map[string]int)
	ages["Alice"] = 32
	ages["Amery"] = 18

	for name, age := range ages {
		fmt.Printf("%s 's age is %d \n", name, age)
	}

	locs := map[string]string{}
	locs["Alice"] = "USA"
	locs["Amery"] = "UK"
	locs["Stan"] = "Canada"

	for name, _ := range locs {
		if ages[name] == 0 {
			fmt.Printf("ages does not contain %s \n", name)
		}
	}

	delete(locs, "Stan")
	delete(locs, "none")

	fmt.Println(locs)

	//charcount()
	wordfreq()
}

//练习 4.8： 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。
func charcount() {
	counts := make(map[string]int) // counts of Unicode characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(r) {
			counts["L"] += 1
		}
		if unicode.IsNumber(r) {
			counts["N"] += 1
		}
		if unicode.IsMark(r) {
			counts["M"] += 1
		}
	}
	fmt.Printf("type \t count\n")
	for c, n := range counts {
		fmt.Printf("%s \t %d \n", c, n)
	}
}

//练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
func wordfreq() {
	counts := make(map[string]int) // counts of Unicode characters
	in := bufio.NewReader(os.Stdin)
	for {
		scanner := bufio.NewScanner(in)
		scanner.Split(bufio.ScanWords)
		// Count the words.
		for scanner.Scan() {
			counts[scanner.Text()] += 1
		}
		_, _, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "system error: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("word \t count\n")
	for c, n := range counts {
		fmt.Printf("%s \t %d \n", c, n)
	}
}
