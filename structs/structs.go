package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(n string) *person {
	p := person{name: n}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	sperson := person{name: "Sean", age: 50}
	fmt.Println(sperson.name)

	sp := &sperson
	fmt.Println(sp.age)

	sp.age = 41
	fmt.Println(sp.age)
}
