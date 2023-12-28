package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{
	// 	"john",
	// 	"doe",
	// }
	// anderson := person{
	// 	firstName: "John",
	// 	lastName:  "Doe",
	// }
	var alex person
	alex.firstName = "John"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
