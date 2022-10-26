package main

import "fmt"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var s1, s2 string
	for _, w := range word1 {
		s1 += w
	}
	for _, w := range word2 {
		s2 += w
	}

	fmt.Println(s1, s2)
	return s1 == s2
}

func main() {
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}
