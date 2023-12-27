package main

import "fmt"

type Person struct {
	name        string
	phoneNumber int
}

func sayName(person Person) string {
	p := person
	return p.name

}

func main() {
	// var is used to declare variable
	// var integer int
	// // var boolean bool
	// // var name string
	// // infered aaignment with theshort assignment operator
	// boolean := true // short assignement for booleans
	// var name = "hello"
	// floating := 2.3
	// // declaring multiple variables on the same line
	// age, lastName := 2, "wade"

	// TYPE SIZES
	// Ints, uints, floats, and complex numbers all have type sizes.

	// int  int8  int16  int32  int64 // whole numbers

	// uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers

	// float32 float64 // decimal numbers

	// complex64 complex128 // imaginary numbers (rare)
	// The size (8, 16, 32, 64, 128, etc) indicates how many bits in memory will be used to store the variable. The default int and uint types are just aliases that refer to their respective 32 or 64 bit sizes depending on the environment of the user.

	// The standard sizes that should be used unless the developer has a specific need are:

	// int
	// uint
	// float64
	// complex128

	// 	FORMATTING STRINGS IN GO
	// Go follows the printf tradition from the C language. In my opinion, string formatting/interpolation in Go is currently less elegant than JavaScript and Python.

	// fmt.Printf - Prints a formatted string to standard out.
	// fmt.Sprintf() - Returns the formatted string
	// EXAMPLES
	// These formatting verbs work with both fmt.Printf and fmt.Sprintf.

	// %V - INTERPOLATE THE DEFAULT REPRESENTATION
	// The %v variant prints the Go syntax representation of a value. You can usually use this if you're unsure what else to use. That said, it's better to use the type-specific variant if you can.

	// s := fmt.Sprintf("I am %v years old", 10)
	// // I am 10 years old

	// s := fmt.Sprintf("I am %v years old", "way too many")
	// // I am way too many years old
	// %S - INTERPOLATE A STRING
	// s := fmt.Sprintf("I am %s years old", "way too many")
	// // I am way too many years old
	// %D - INTERPOLATE AN INTEGER IN DECIMAL FORM
	// s := fmt.Sprintf("I am %d years old", 10)
	// // I am 10 years old
	// %F - INTERPOLATE A DECIMAL
	// s := fmt.Sprintf("I am %f years old", 10.523)
	// // I am 10.523000 years old

	// // The ".2" rounds the number to 2 decimal places
	// s := fmt.Sprintf("I am %.2f years old", 10.523)
	// // I am 10.53 years old

	// CONDITIONALS
	// if statements in Go don't use parentheses around the condition:

	// if height > 4 {
	//     fmt.Println("You are tall enough!")
	// }
	// else if and else are supported as you would expect:

	// if height > 6 {
	//     fmt.Println("You are super tall!")
	// } else if height > 4 {
	//     fmt.Println("You are tall enough!")
	// } else {
	//     fmt.Println("You are not tall enough!")
	// }

	// 	THE INITIAL STATEMENT OF AN IF BLOCK
	// An if conditional can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if body.

	// if INITIAL_STATEMENT; CONDITION {
	// }
	// WHY WOULD I USE THIS?
	// This is just some syntactic sugar that Go offers to shorten up code in some cases. For example, instead of writing:

	// length := getLength(email)
	// if length < 1 {
	//     fmt.Println("Email is invalid")
	// }
	// We can do:

	// if length := getLength(email); length < 1 {
	//     fmt.Println("Email is invalid")
	// }

	// fmt.Println(integer, boolean, name, floating, age, lastName)
	// structs
	func chaeck (){
		fmt.printf("hello world")
	}

	person1 := Person{
		name: "wade", phoneNumber: 123,
	}
	fmt.Printf(sayName(person1))

	fmt.Printf("hello world")
}
