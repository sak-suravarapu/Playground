package main

import (
	"fmt"
	"std/hello_world/motd/message"
)

func main() {
	message := message.Greeting("Ram", "Hello")
	fmt.Println(message)
	fmt.Println(greeting1("Robert", "Bye"))

}

func greeting1(name, message string) (response string) {
	response = fmt.Sprintf("%s, %s", message, name)
	return
}
