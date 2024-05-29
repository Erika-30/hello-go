// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello, World!")
// }

package main

import "fmt"

func main() {
	fmt.Printf("%v\n", 123)       // Prints 123
	fmt.Printf("%#v\n", 123)      // Prints 123 in Go-syntax representation
	fmt.Printf("%T\n", 123)       // Prints int
	fmt.Printf("%%\n")            // Prints %
	fmt.Printf("%t\n", true)      // Prints true
	fmt.Printf("%b\n", 123)       // Prints 1111011 in binary
	fmt.Printf("%c\n", 123)       // Prints {
	fmt.Printf("%d\n", 123)       // Prints 123 in decimal
	fmt.Printf("%o\n", 123)       // Prints 173 in octal
	fmt.Printf("%x\n", 123)       // Prints 7b in hexadecimal (lowercase)
	fmt.Printf("%X\n", 123)       // Prints 7B in hexadecimal (uppercase)
	fmt.Printf("%e\n", 123.456)   // Prints 1.234560e+02 in scientific notation
	fmt.Printf("%f\n", 123.456)   // Prints 123.456000 in floating point
	fmt.Printf("%g\n", 123.456)   // Prints 123.456 in compact notation
	fmt.Printf("%s\n", "hello")   // Prints hello
	fmt.Printf("%q\n", "hello")   // Prints "hello" in quotes

	// Assign the mainFunc function to a variable
	funcPtr := mainFunc
	// Print the address of the variable that holds the function
	fmt.Printf("%p\n", &funcPtr)  // Prints the memory address of funcPtr
}

func mainFunc() {
	fmt.Println("Hello, World!")
}

