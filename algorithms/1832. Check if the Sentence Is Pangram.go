package main

func checkIfPangram(sentence string) bool {
	var m = make(map[rune]bool)
	for _, c := range sentence {
		m[c] = true
	}
	return len(m) == 26
}
