package main

import (
	"fmt"
)

func main() {
	/*
		第一段go程序
	*/
	mahesh := "super man"
	fmt.Println("hello ,go")
	fmt.Println(mahesh)
	var age int
	age = 35
	fmt.Println(age)

	var money = 36.5
	var salary float64
	salary = 88.555
	fmt.Println(money)
	fmt.Println(salary)

	var name string
	name = "karel"
	fmt.Println(name)

	address := "陕西，西安"
	fmt.Println(address)

	var hasEaten bool
	if len(address) > 0 {
		hasEaten = true
	}
	if hasEaten {
		fmt.Println("吃过饭了，那就准备睡觉了啊")
	}

	var books [5]string
	books[0] = "深入理解计算机系统"
	books[1] = "编译原理"
	fmt.Println(books)

	i := 0
	for i = 0; i < len(books); i++ {
		fmt.Println(books[i])
	}
}
