package main

import "fmt"

type rect struct {
	width, height uint
}

func (r *rect) area() uint {
	return r.width * r.height
}

func (r rect) perimeter() uint {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 18, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
