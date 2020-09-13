package main

import (
	"fmt"
)

type person struct {
	// when providing field names, the first
	// letter being lowercase denotates
	// fields that can only be r-w inside
	// the package where the struct is written in
	name string
	age  int
}

// Creature struct exportable
type Creature struct {
	// Uppercase letters denotates exportable fields
	// that can be used elsewhere
	Name string
	Type string
}

// CreatureWithSecrets you can also do mixed structs
// with exported and non-exported fields
type CreatureWithSecrets struct {
	Name string
	Type string

	sekret string
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

	// note if the instantiation of the struct
	// will include every field then this method
	// of declaring is correct.
	// Otherwise it may yield compile errors.

	// This method of instantiating structs
	// should only be done in VERY small structs
	// as it can get confusing fast.
	fmt.Println(person{"Bob", 20})

	// you can also explicitly define structs
	// this is the most preferred method of instantiating
	// structs as you don't have to provide nil-like values
	// in the short form.
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

	// it is also possible to declare inline structs
	// thus creating anonymous structs and avoiding naming
	// structs that would only be used for 1 purpose
	// and 1 purpose alone.

	// this usually occurs during test suites
	digitalWowcean := struct {
		Name string
		Type string
	}{Name: "mammasharkdododaddysharkdodo", Type: "shark"}

	fmt.Println(digitalWowcean.Name, digitalWowcean.Type)

	fmt.Println(CreatureWithSecrets{Name: "hello", sekret: "sekret"})
}
