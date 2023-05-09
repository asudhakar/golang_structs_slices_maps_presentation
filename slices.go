package main

import "fmt"

func main() {
	//EG1START OMIT
	var numbers []int // An empty slice of integers
	//EG1END OMIT
	//EG2START OMIT
	numbers = []int{1, 2, 3, 4, 5}
	//EG2END OMIT
	//EG3START OMIT
	numbers = append(numbers, 6)
	//EG3END OMIT
	fmt.Println(numbers[2]) // Prints the third element
	numbers[2] = 10         // Changes the value of the third element
	fmt.Println(numbers[2]) // Prints the third element
	// Lets Print all elements // HL
	fmt.Println(numbers)

}
