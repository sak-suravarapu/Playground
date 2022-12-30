package main

import "fmt"

func main() {

	ages := map[string]int{}
	ages["Ram"] = 7
	ages["Rahim"] = 6
	ages["Robert"] = 9
	ages["Lakshmi"] = 10
	for name, age := range ages {
		switch age {
		case 1, 3, 5, 7, 9:
			fmt.Println(name, "Age is odd numbers less than 10")
		case 2, 4, 6, 8:
			fmt.Println(name, "Age is even numbers less than 10")
		default:
			fmt.Println(fmt.Sprintf("There is nothing special about %s's age", name))
		}
	}
	for i := 1; i <= 10; i++ {
		fmt.Println("We are counting:", i)
	}

	a := 0
	for a < 10 {
		if a%2 == 0 {
			fmt.Println("We are counting:", a)
			a++
			continue
		} else if a == 5 {
			break
		}
	}
}
