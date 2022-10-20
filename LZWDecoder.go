package main

// Returns a dictionary mapping values 0x00 - 0xFF
// to their respective ASCII counterparts
func LZWGetInitialDict() map[int]string {
	dict := make(map[int]string)
	for i := 0; i < 256; i++ {
		dict[i] = string(byte(i))
	}
	return dict
}

// Performs LZW Decoding on a list of integers
func LZWDecode(codes []int) string {
	dict := LZWGetInitialDict()
	// Manually decode first code
	w := dict[int(codes[0])]
	out := w

	for i, code := range codes {
		// Skip first code
		if i == 0 {
			continue
		}
		// check to for need of dictionary reset
		if len(dict) == 4096 {
			dict = LZWGetInitialDict()
		}
		entry := ""
		// decode current code
		if val, found := dict[code]; found {
			entry = val
		} else if code == len(dict) {
			entry = w + w[0:1]
		}
		// emit data
		out += entry

		// add code to dictionary
		dict[len(dict)] = w + entry[0:1]
		w = entry

	}
	return out
}
