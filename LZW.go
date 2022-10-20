package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for Args
	if !(len(os.Args) > 1) {
		fmt.Println("Usage: go run . [input file path] [output file path]")
		return
	}
	infile := os.Args[1]
	outfile := os.Args[2]

	// Read as 12bit
	codes := ReadFileAs12bit(infile)

	// Decode
	out := LZWDecode(codes)

	// Output
	WriteStringToFile(outfile, out)
}
