package main

func LZWGetInitialDict() map[int]string {
	dict := make(map[int]string)
	for i := 0; i < 256; i++ {
		dict[i] = string(rune(i))
	}
	return dict
}

func LZWDecode(codes []uint) string {
	out := ""
	previous := ""
	dict := LZWGetInitialDict()
	for i := 0; i < len(codes); i++ {
		symbol := int(codes[i])
		if val, found := dict[symbol]; found {
			out += val
			dict[len(dict)] = previous + string(rune(symbol))
			previous = val
		} else {
			V := previous + string(rune(symbol))
			dict[len(dict)] = V
			out += val
			previous = V
		}
	}
	return out
}
