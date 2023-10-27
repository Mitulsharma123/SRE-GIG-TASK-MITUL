package main

import (
	"math/rand"
	"fmt"
)

func main(){
	n := rand.Intn(100)	//generate a random number between 1-100
	fun1(n)	// as per point 1 from problem statement 
	fun2(n)	// as per point 2 from problem statement
	fun3(n)	// as per point 3 from problem statement
}

func fun1(num int) {
	fmt.Printf("Generated random number is: %d\n",num)
		if num > 50 {
			fmt.Println("It's closer to 100")
		} 
		if num < 50 {
			fmt.Println("It's closer to 0")
		} 	
}

func fun2(num int) {
	fmt.Printf("Generated random number is: %d\n",num)
		if num == 50 {
			fmt.Println("It's 50!")
		} 
		if num < 50 {
			fmt.Println("It's closer to 0", num)
	}
}

func fun3(num int) {
	fmt.Printf("Generated random number is: %d\n",num)
		if num > 50 && num%2 == 0{	// if remainder is zero, it is even number 
			fmt.Println("It's closer to 100, and it's even!", num)
		}  
		if num < 50 {
			fmt.Println("It's closer to 0", num)
		}
}