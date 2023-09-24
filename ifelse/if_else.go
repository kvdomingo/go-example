package main

import "fmt"

func main() {
	var i = 9
	i += 1

	if i%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
