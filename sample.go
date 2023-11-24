package main

import "fmt"

func main() {
	num := 123
	fmt.Print("Hello, world", num)
	fmt.Println("Hello, world", num)
	fmt.Printf("Hello, world%d", num)

	const pi = 3.14
}
