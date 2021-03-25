package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf(comma("4471444755"))
	fmt.Println()
	fmt.Printf(comma2("4471444755"))
	fmt.Println()
	fmt.Printf(basename("a/b/c.go"))
	fmt.Println()
	fmt.Printf(basename2("a/b/c.go"))
	fmt.Println()
	b := equalNotOrder("hello, world", "worldhell o,")
	b1 := equalNotOrder("1112", "2123")
	fmt.Printf("%t", b)
	fmt.Println()
	fmt.Printf("%t", b1)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//5874.41
//todo
func commaFloat(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	tmp := s[:n-3]
	if strings.Contains(tmp, ".") {
		return commaFloat(tmp) + "," + s[n-3:]
	} else {
		return commaFloat(tmp) + "," + s[n-3:]
	}

}

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	j := 0
	for i := n; i > 0; i-- {
		if j == 3 {
			buf.WriteString(",")
			j = 0
		}
		buf.WriteString(s[i-1 : i])
		j += 1
	}
	sb := buf.String()

	var result string
	for _, v := range sb {
		result = string(v) + result
	}
	return result
}

func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func equalNotOrder(s1, s2 string) bool {
	for _, v := range s1 {
		if !exist(v, s2) {
			return false
		}
	}
	for _, v := range s2 {
		if !exist(v, s1) {
			return false
		}
	}
	return true
}

func exist(v int32, s2 string) bool {
	for _, v2 := range s2 {
		if v2 == v {
			return true
		}
	}
	return false
}
