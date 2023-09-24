package main

import "fmt"

type jedi struct {
	name  string
	color string
	age   uint
}

func newJedi(name string) *jedi {
	j := jedi{name: name}
	j.age = 13
	j.color = "blue"
	return &j
}

func main() {
	fmt.Println(jedi{"Luke", "blue", 13})
	fmt.Println(&jedi{"Anakin", "red", 32})
	fmt.Println(newJedi("Obi-Wan"))

	s := jedi{name: "Anakin", color: "red", age: 32}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sith := struct {
		name    string
		isAlive bool
	}{
		"Darth Vader",
		false,
	}
	fmt.Println(sith)
}
