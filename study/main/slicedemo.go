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
//?????? 4.3??? ??????reverse?????????????????????????????????slice???
func reverse2(ptr *[6]int) {
	for i, j := 0, len(ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

//?????? 4.5??? ????????????????????????????????????[]string???????????????????????????????????????
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

//?????? 4.6??? ????????????????????????????????????UTF-8?????????[]byte?????????slice???????????????????????????unicode.IsSpace??????????????????????????????
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

//?????? 4.7??? ??????reverse????????????????????????UTF-8?????????[]byte?????????????????????????????????????????????
func reverse3(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
