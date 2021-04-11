package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	arr := [32]int{31: 9}
	println(arr[31]) //prints 9
	zero1(arr)
	println(arr[31]) //prints 9
	zero(&arr)
	println(arr[31]) //prints 0

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

var pc [256]byte

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func zero(ptr *[32]int) {
	*ptr = [32]int{}
}

/**
On default , function use a copy of parameter
*/
func zero1(ptr [32]int) {
	ptr = [32]int{}
}
