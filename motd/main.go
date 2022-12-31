package main

import (
	"bufio"
	"fmt"
	"hello_world/motd/message"
	"os"
	"strings"
)

func main() {

	f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error: Unable to open /etc/motd")
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)
	fmt.Println("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	m := message.Greeting(name, phrase)
	fmt.Println(m)
	_, err = f.Write([]byte(m))
	//err := ioutil.WriteFile("/etc/motd", []byte(m), 0644)
	if err != nil {
		fmt.Println("Error: Failed to write to /etc/motd")
		os.Exit(1)
	}
}
