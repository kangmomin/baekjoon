package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)

	if input%4 == 0 && input%100 != 0 || input%400 == 0 {
		fmt.Print(1)
		return
	}

	fmt.Print(0)
}
