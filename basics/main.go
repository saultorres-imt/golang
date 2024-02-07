package main

import "fmt"

func main() {
	useFmt()
}

/*
The 'fmt' package in Go is one of the most essential and frequently used libraries in the Go programming language.
It provides a wide range of functionalities for formatting and printing of output, as well as taking input.
Key functions include:
  - fmt.Print(), fmt.Println(), fmt.Printf(): Used for printing output to the console. Print() outputs text as is,
    Println() adds a newline at the end, and Printf() allows for formatted output using verbs like %v, %d, %s, etc.
  - fmt.Scanf(), fmt.Scanln(), fmt.Scan(): These functions are used for taking input from the user. Scanf() reads
    formatted input, Scanln() reads input until a newline is encountered, and Scan() reads input separated by spaces.
  - fmt.Sprintf(): Similar to Printf(), but instead of printing to the console, it returns the formatted string.
  - fmt.Errorf(): Used for creating error messages with formatted text, handy for error handling with descriptive messages.

https://pkg.go.dev/fmt
*/
func useFmt() {
	// Adds a newline at the end of a print
	fmt.Println("Hello World 😁")

	// Common Format Specifiers in fmt.Printf:
	// - %v: The value in a default format. When printing structs, the plus flag (%+v) adds field names.
	// - %T: A Go-syntax representation of the type of the value.
	// - %d: Base 10 integer formatting.
	// - %b: Base 2 (binary) integer formatting.
	// - %x, %X: Base 16 (hexadecimal), lowercase and uppercase formatting.
	// - %f: Floating point number formatting.
	// - %s: Basic string formatting.
	// - %q: Quoted string formatting (adds double quotes around strings).
	// - %p: Pointer address formatting (prints the memory address of a value).
	fmt.Printf("Printing string %q\nType %T\nBinary%b\n", "Hello", 12, 5)
}
