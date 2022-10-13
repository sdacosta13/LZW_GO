package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileAs12bit(filename string) []uint {
	// Read in bytes and handle errors
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

		// Bit manipulation to shift the binary to the correct order for addition
		leftInt := (uint(byteA) << 4) + uint(byteB>>4)
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
