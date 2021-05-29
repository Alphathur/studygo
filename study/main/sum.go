package main

import (
	"fmt"
	"strings"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	min := vals[0]
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}

func main() {
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	fmt.Println(max(values...))
	fmt.Println(min(values...))

	s := strings.Join([]string{"a", "b", "c"}, "0")
	fmt.Println(s)
}
