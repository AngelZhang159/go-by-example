package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [...]int{1, 2}
	fmt.Println("dcl:", b)

	c := [...]int{1, 3, 5, 6}
	fmt.Println("dcl:", c)

	d := [...]int{100, 3: 400, 500}
	fmt.Println("idx:", d)

	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println("2d:", twoD)

}
