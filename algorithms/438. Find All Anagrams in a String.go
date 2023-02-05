package main

func findAnagrams(s string, p string) []int {
	var res []int
	// 1. 如果s的长度小于p的长度，直接返回
	if len(s) < len(p) {
		return res
	}
	// 2. 统计p中每个字符出现的次数
	var pMap [26]int
	// 3. 统计s中每个字符出现的次数
	for _, v := range p {
		pMap[v-'a']++
	}
	// 4. 统计s中每个字符出现的次数
	var sMap [26]int
	// 5. 统计s中前len(p)个字符出现的次数
	for i := 0; i < len(p); i++ {
		sMap[s[i]-'a']++
	}
	// 6. 如果s中前len(p)个字符出现的次数和p中每个字符出现的次数相同，说明s中前len(p)个字符是p的字母异位词
	if pMap == sMap {
		res = append(res, 0)
	}
	// 7. 统计s中后面的字符出现的次数
	for i := len(p); i < len(s); i++ {
		// 8. 统计s中后面的字符出现的次数
		sMap[s[i]-'a']++
		// 9. 统计s中前面的字符出现的次数
		sMap[s[i-len(p)]-'a']--
		// 10. 如果s中前len(p)个字符出现的次数和p中每个字符出现的次数相同，说明s中前len(p)个字符是p的字母异位词
		if pMap == sMap {
			// 11. 记录s中前len(p)个字符的起始位置
			res = append(res, i-len(p)+1)
		}
	}
	return res
}
