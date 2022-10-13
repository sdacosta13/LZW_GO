package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for Args
	if !(len(os.Args) > 1) {
		fmt.Println("A file must be specified")
		return
	}
	filename := os.Args[1]
	codes := ReadFileAs12bit(filename)
	out := LZWDecode(codes)
	fmt.Print(out)
}
