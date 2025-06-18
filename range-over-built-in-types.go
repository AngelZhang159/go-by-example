package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for s2, s3 := range kvs {
		fmt.Printf("%s -> %s\n", s2, s3)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, i2 := range "go" {
		fmt.Println(i, i2)
	}
}
