package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, el := range arr {
		sum += el
	}
	fmt.Println(sum)

	kvs := map[string]string{
		"name":  "Anakin",
		"color": "red",
	}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for index, char := range "golang" {
		fmt.Println(index, char)
	}
}
