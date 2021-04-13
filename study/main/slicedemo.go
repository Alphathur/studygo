package main

import (
	"fmt"
	"unicode"
)

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months)
	fmt.Println(months[2])
	fmt.Println(cap(months))
	fmt.Println(len(months))

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	fmt.Println(len(summer)) //3
	fmt.Println(cap(summer)) //7
	//fmt.Println(summer[3]) //visit ,panic: runtime error: index out of range [3] with length 3
	//fmt.Println(summer[:20]) //extend, panic: runtime error: slice bounds out of range [:20] with capacity 7
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"
	fmt.Println(cap(summer))

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`

	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data2)) // `["one" "three"]`
	fmt.Printf("%q\n", data2)            // `["one" "three" "three"]`

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	a1 := [6]int{0, 1, 2, 3, 4, 5}
	reverse2(&a1)

	a2 := []string{"a", "a", "b", "b", "c"}
	a2 = nonduplicate(a2)
	fmt.Println(a2)

	s := "Hello, friends"
	b := nonempty3(s)
	fmt.Println(string(b))

	s1 := "Hello,     friends"
	b1 := nonsamespace([]rune(s1))
	fmt.Println(string(b1))

	bs := []byte(s)
	reverse3(bs)
	fmt.Println(string(bs))
}

func deepEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//use arr and pointer to reverse
//练习 4.3： 重写reverse函数，使用数组指针代替slice。
func reverse2(ptr *[6]int) {
	for i, j := 0, len(ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

//练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func nonduplicate(strings []string) []string {
	var out []string
	for i := range strings {
		if i == 0 {
			out = append(out, strings[i])
			continue
		}
		if strings[i] != strings[i-1] {
			out = append(out, strings[i])
		}
	}
	return out
}

func nonempty3(str string) []int32 {
	var out []int32
	for _, s := range str {
		if !unicode.IsSpace(s) {
			out = append(out, s)
		}
	}
	return out
}

//练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func nonsamespace(r []rune) []rune {
	var out []rune
	for i := range r {
		if i == 0 {
			out = append(out, r[i])
			continue
		}
		if unicode.IsSpace(r[i]) && unicode.IsSpace(r[i-1]) {
			continue
		} else {
			out = append(out, r[i])
		}
	}
	return out
}

//练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
func reverse3(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
