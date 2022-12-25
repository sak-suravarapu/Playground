package main

import "fmt"

func main() {
	fmt.Println("Simple String")
	fmt.Println(`
	This is a multi-line String, that can also contain "quotes".
	String
	`)
	fmt.Println("😀")
	fmt.Println("\u2772")
	fmt.Println('L')
}
