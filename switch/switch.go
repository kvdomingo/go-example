package main

import "fmt"

func main() {
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("wtf is this %T", t)
		}
	}

	whatAmI(true)
	whatAmI(23)
	whatAmI("hello there")
}
