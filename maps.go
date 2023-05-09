package main

import "fmt"

func main() {
	//EG1START OMIT
	m := make(map[string]int) // An empty map with string keys and integer values
	//EG1END OMIT
	//EG2START OMIT
	m["apple"] = 1
	m["banana"] = 2
	//EG2END OMIT
	//EG3START OMIT
	fmt.Println(m["apple"]) // Prints the value associated with the key "apple"
	//EG3END OMIT
	// Lets Print all data // HL
	fmt.Printf("%v \n", m)
	//EG4START OMIT
	delete(m, "banana") // Deletes the key-value pair with the key "banana"
	//EG4END OMIT
	// Lets see the data after deleting banana // HL
	fmt.Printf("%v", m)
}
