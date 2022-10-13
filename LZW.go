package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Check for Args
	if !(len(os.Args) > 1) {
		fmt.Println("A file must be specified")
		return
	}
	dat, err := os.ReadFile(os.Args[1])
	check(err)
	fmt.Print(string(dat))
}
