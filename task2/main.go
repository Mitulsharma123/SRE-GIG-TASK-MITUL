package main

import (
	"math/rand"
	"fmt"
)

func main(){
	n := rand.Intn(100)
	fun1(n)
	fun2(n)
	fun3(n)
}

func fun1(num int) {
	if num > 50 {
		fmt.Println("It's closer to 100")
	} else if num < 50 {
		fmt.Println("It's closer to 0")
	} else {
		fmt.Println("Generated number:", num)
	}
}

func fun2(num int) {
	if num == 50 {
		fmt.Println("It's 50!")
	} else if num < 50 {
		fmt.Println("It's closer to 0")
	}else {
		fmt.Println("Generated number:", num)
	}
}

func fun3(num int) {
	if num > 50 && num%2 == 0{
		fmt.Println("It's closer to 100, and it's even!")
	} else if num < 50 {
		fmt.Println("It's closer to 0")
	} else {
		fmt.Println("Generated number:", num)
	}
}