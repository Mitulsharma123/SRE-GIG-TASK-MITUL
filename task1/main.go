package main

import "fmt"

func main(){
	var Menu [2]string	
	
	arr := [5]string{"apple", "oranges", "banana", "papaya", "peach"}
	
	Menu[0]="hamburger"
	Menu[1]="salad"

	for _, item := range Menu{
		fmt.Printf("List of food item found in menu: %s\n", item)
		
	}

	fmt.Println()

	for index, val := range arr{
		fmt.Printf("This is %s and its index in the array is %d \n",val, index)
	}
}

