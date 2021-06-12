package main

import "fmt"

type person struct {
	name string
	age  int
}

func (s person) Read() string {
	return s.name
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	stars := person{"Bob", 20}
	fmt.Println(stars.Read())

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
