package main

import "os"

func main() {

	//panics are errors that should never happen and will crash the app

	//panic("problem!!")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
