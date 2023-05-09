package main

import "fmt"

func main() {
	//EG1START OMIT
	type Person struct {
		Name string
		Age  int
	}
	//EG1END OMIT
	//EG2START OMIT
	p := Person{Name: "John Doe", Age: 30}
	//EG2END OMIT
	//EG3START OMIT
	fmt.Println(p.Name, p.Age)
	//EG3END OMIT
}
