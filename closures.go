package main

import "fmt"

func intSeq() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func main() {
	plus := intSeq()

	ret := plus()
	fmt.Println(ret)
	ret = plus()
	ret = plus()
	ret = plus()
	ret = plus()
	fmt.Println(ret)

}
