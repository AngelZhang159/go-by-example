package main

import (
	"fmt"
	"os"
)

func main() {

	//defer is used to call a func later in the program, mainly used for cleanup

	f := createFile("tmp", "defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func writeFile(file *os.File) {
	fmt.Println("writing")
	_, err := fmt.Fprintln(file, "data")
	if err != nil {
		panic(err)
	}
}

func closeFile(file *os.File) {
	fmt.Println("closing")
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func createFile(dir string, file string) *os.File {
	fmt.Println("creating")
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(dir, 0777)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(dir + "/" + file)
	if err != nil {
		panic(err)
	}
	return f
}
