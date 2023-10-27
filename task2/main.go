package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)


func main(){
	var max big.Int
	max.SetInt64(100)
	n, err := rand.Int(rand.Reader, &max)	//generate a random number between 1-100
	if err != nil{
		panic(err)
	}
	random := n.Int64()
	//fmt.Println(random)

	fun1(random)	// as per point 1 from problem statement 
	fun2(random)	// as per point 2 from problem statement
	fun3(random)	// as per point 3 from problem statement
}

func fun1(num int64) {
	fmt.Printf("Generated random number is: %d\n",num)
		if num > 50 {
			fmt.Println("It's closer to 100")
		} 
		if num < 50 {
			fmt.Println("It's closer to 0")
		} 	
}

func fun2(num int64) {
	fmt.Printf("Generated random number is: %d\n",num)
		if num == 50 {
			fmt.Println("It's 50!")
		} 
		if num < 50 {
			fmt.Println("It's closer to 0", num)
	}
}

func fun3(num int64) {
	fmt.Printf("Generated random number is: %d\n",num)
		if num > 50 && num%2 == 0{	// if remainder is zero, it is even number 
			fmt.Println("It's closer to 100, and it's even!", num)
		}  
		if num < 50 {
			fmt.Println("It's closer to 0", num)
		}
}