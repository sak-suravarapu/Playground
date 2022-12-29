package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Ram"] = 7
	switch {
	case ages["Ram"] < 18:
		fmt.Println("Ram cannot vote")
	case ages["Ram"] < 67:
		fmt.Println("Ram is not of retirement age")
	default:
		fmt.Println("Ram is of retirement age")
	}
	switch ages["Ram"] {
	case 1, 3, 5, 7, 9:
		fmt.Println("Age is odd numbers less than 10")
	case 2, 4, 6, 8:
		fmt.Println("Age is even numbers less than 10")
	}
	if ages["Ram"] < 18 {
		fmt.Println("Ram cannot vote")
	} else if ages["Ram"] < 67 {
		fmt.Println("Ram is not of retirement age")
	} else {
		fmt.Println("Ram is of retirement age")
	}
}
