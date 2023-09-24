package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s)

	s = make([]string, 3)
	fmt.Println(s)

	s[0] = "anakin"
	s[1] = "obi-wan"
	s[2] = "luke"
	fmt.Println(s)
}
