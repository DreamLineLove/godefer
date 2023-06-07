package main

import "fmt"

func main() {
	defer fmt.Println("the functions that have defer keyword before their names, are added onto a stack.")
	defer fmt.Println("when the enclosing function returns, the stack is called.")
	defer fmt.Println("being a stack, the last function in, is the first one to be executed.")
	fmt.Println("Hello World")
}
