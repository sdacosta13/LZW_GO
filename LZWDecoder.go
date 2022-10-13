package main

func LZWGetInitialDict() [][]string {
	entries := [][]string{}
	for i := 0; i < 256; i++ {
		row := []string{"", "", string(rune(i))}
		entries = append(entries, row)
	}
	return entries
}

func LZWDecode() {
	entries := LZWGetInitialDict()

}
