package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads data in the 12bit format specified
func ReadFileAs12bit(filename string) []int {
	// Read in bytes and handle errors
	data, err := os.ReadFile(filename)
	check(err)

	//account for odd number of codes
	effectiveLength := len(data)
	if len(data)%3 != 0 {
		effectiveLength -= 2
	}
	codes := []int{}
	for i := 0; i < effectiveLength; i += 3 {
		// examine three bytes at a time and extract 2 numbers
		byteA := data[i+0]
		byteB := data[i+1]
		byteC := data[i+2]

		// Bit manipulation to shift the binary to the correct order for addition
		leftInt := (int(byteA) << 4) + int(byteB>>4)
		rightInt := (int(byteB<<4) << 4) + int(byteC)

		codes = append(codes, leftInt, rightInt)
	}
	if len(data)%3 != 0 {
		// calculates final byte if codes are uneven
		leftByte := data[len(data)-2]
		rightByte := data[len(data)-1]
		finalInt := (int(leftByte<<4) << 4) + int(rightByte)
		codes = append(codes, finalInt)
	}
	return codes
}

// Creates a file at the path and writes the data
func WriteStringToFile(path string, data string) {
	f, e := os.Create(path)
	check(e)
	defer f.Close()
	_, e = f.WriteString(data)
	check(e)
}
