package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{num: 5},
		str:  "something",
	}

	fmt.Println(co)
	fmt.Println(co.base.num)
	fmt.Println(co.base.describe())

	var d describer = co
	fmt.Println("describer:", d.describe())

	var bas describer = co.base
	fmt.Println("describer:", bas.describe())
}

type describer interface {
	describe() string
}
