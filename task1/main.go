package main

import "fmt"

func main(){
	var Menu [2]string	//declare an array 

	arr := [5]string{"apple", "oranges", "banana", "papaya", "peach"} // declare an array with some values
	
	Menu[0]="hamburger"	// add values to array Menu at index 0
	Menu[1]="salad"		/// add values to array Menu at index 1

	// iterate over array menu with for loop
	// each iteration will print the values inside the array menu
	for _, item := range Menu{
		fmt.Printf("List of food item found in menu: %s\n", item)
		
	}

	fmt.Println()

	// iterate over array arr with for loop
	// each iteration will print the values inside the arr menu along with the index
	for index, val := range arr{
		fmt.Printf("This is %s and its index in the array is %d \n",val, index)
	}
}

