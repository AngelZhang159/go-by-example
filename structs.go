package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{
		name: "Alice",
		age:  30,
	})

	fmt.Println(person{
		name: "Fred",
	})

	fmt.Println(&person{
		name: "Ann",
		age:  40,
	})

	s := person{
		name: "Sean",
		age:  50,
	}
	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)

}
