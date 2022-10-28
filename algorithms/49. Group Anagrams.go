package main

import "strconv"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		key := getKey(str)
		m[key] = append(m[key], str)
	}

	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func getKey(str string) string {
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}

	var result string
	for i := 0; i < 26; i++ {
		if m[byte(i+'a')] > 0 {
			result += string(byte(i+'a')) + strconv.Itoa(m[byte(i+'a')])
		}
	}

	return result
}
