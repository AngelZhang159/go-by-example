package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	arr := []int{1, 2, 3, 4, 5, 6}
	sum(arr...)
}
