package main

func LZWGetInitialDict() map[int]string {
	dict := make(map[int]string)
	for i := 0; i < 256; i++ {
		dict[i] = string(rune(i))
	}
	return dict
}

func LZWDecode(codes []uint) string {
	dict := LZWGetInitialDict()
	w := dict[int(codes[0])]
	out := w
	for i, code := range codes {
		if i == 0 {
			continue
		}
		entry := ""
		if val, found := dict[int(code)]; found {
			entry = val
		} else if int(code) == len(dict) {
			entry = w + w[0:1]
		}
		out += entry
		dict[len(dict)] = w + entry[0:1]
		w = entry
	}
	return out
}
