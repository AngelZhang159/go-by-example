package main

import "fmt"

type rect struct {
	width  int
	height int
}

// In this example, both func are the same bc they are not modifying the og struct passed
// If you pass a pointer, then any change you make in the func reflects in the original struct
// If you pass the value itself it makes a copy and the original is not affected
func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() any {
	return 2 * (r.width + r.height)
}

func main() {

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perim())

	rp := &r

	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter: ", rp.perim())

}
