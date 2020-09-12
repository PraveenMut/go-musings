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
	// output with a values passed in
	// the compiler automatically the structure
	// of the structs passed on the positionals args
	fmt.Println(person{"Bob", 20})

	// you can also explicitly define structs
	fmt.Println(person{name: "Alice", age: 30})

	// any values that are not passed in
	// to additional/other fields are
	// zero-valued (nulled)
	fmt.Println(person{name: "Fred"})

	// adding the & operator passes the
	// memory address which becomes deferenced
	fmt.Println(&person{name: "Ann", age: 40})

	// this where the function passed as values
	// works
	fmt.Println(newPerson("Jon"))

	// sperson is passed as a normal struct
	// but then you can access individual
	// fields of structs by using the .
	// operator
	sperson := person{name: "Sean", age: 50}
	fmt.Println(sperson.name)

	// you can also use pointers by using
	// the & operator as the . operator
	// automatically dereferences the memory address
	sp := &sperson
	fmt.Println(sp.age)

	// you can also directly edit the pointer of
	// the structure and edit values
	// note that this actually changes the struct itself
	// as it passed by reference AND not by value
	sp.age = 41
	fmt.Println(sp.age)
}
