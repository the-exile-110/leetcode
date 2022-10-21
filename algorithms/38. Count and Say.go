package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	var res string
	var count int
	for i := 0; i < len(s); i++ {
		count++
		if i == len(s)-1 || s[i] != s[i+1] {
			res += strconv.Itoa(count) + string(s[i])
			count = 0
		}
	}
	return res
}
