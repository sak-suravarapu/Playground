package main

import "fmt"

func main() {
	names := [3]string{"ram", "robert", "rahim"}
	var names1 [3]string
	names1[0] = "ram"
	names1[1] = "robert"
	names1[2] = "rahim"
	fmt.Println(names)
	fmt.Println(names1)
	fmt.Println("names1 index 2 is nil:", names1[2] == "")
	namessp := []string{}
	namessp = append(namessp, "ram")
	namessp = append(namessp, "rahim", "robert")
	fmt.Println(namessp)
	namesmk := make([]string, 4)
	namesmk[0] = "hello"
	namesmk[1] = "this"
	namesmk[2] = "is"
	namesmk[3] = "ram"
	//namesmk[4] = "rahim"
	//namesmk[5] = "robert"
	fmt.Println(namesmk)

}
