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

func ReadFileAs12bit(filename string) []uint {
	data, err := os.ReadFile(filename)
	check(err)

	effectiveLength := len(data)
	if len(data)%3 != 0 {
		//account for odd number of codes
		effectiveLength -= 2
	}
	codes := []uint{}
	for i := 0; i < effectiveLength; i += 3 {
		// examine three bytes at a time and extract 2 numbers
		byteA := data[i+0]
		byteB := data[i+1]
		byteC := data[i+2]

		// for lefthand
		// first byte shifted 4 to the left to the top of the 12bit int
		// second byte shifted 4 to the right to the bottom of the 12bit int
		//
		//	00000100 10000001 (72) becomes
		//  0000 0100 0000
		// +     0000 1000
		leftInt := (uint(byteA) << 4) + uint(byteB>>4)

		// for righthand
		// first byte shifted left by 4, extended, then shifted by 4 left again
		// second byte simply added
		//  10000001 11110100 (500) becomes
		//  0001 0000 0000
		// +     1111 0100
		rightInt := (uint(byteB<<4) << 4) + uint(byteC)
		codes = append(codes, leftInt, rightInt)
	}
	if len(data)%3 != 0 {
		// calculates final byte if codes are uneven
		leftByte := data[len(data)-2]
		rightByte := data[len(data)-1]
		finalInt := (uint(leftByte<<4) << 4) + uint(rightByte)
		codes = append(codes, finalInt)
	}
	return codes
}

func main() {
	// Check for Args
	if !(len(os.Args) > 1) {
		fmt.Println("A file must be specified")
		return
	}
	filename := os.Args[1]
	codes := ReadFileAs12bit(filename)
	fmt.Print(codes)
}
