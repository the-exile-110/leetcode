package main

func maxLength(arr []string) int {
	var result int
	var dfs func(int, string)
	dfs = func(index int, s string) {
		if len(s) > result {
			result = len(s)
		}
		for i := index; i < len(arr); i++ {
			if hasDuplicate(s + arr[i]) {
				continue
			}
			dfs(i+1, s+arr[i])
		}
	}
	dfs(0, "")
	return result
}

func hasDuplicate(s string) bool {
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if m[s[i]] {
			return true
		}
		m[s[i]] = true
	}
	return false
}
