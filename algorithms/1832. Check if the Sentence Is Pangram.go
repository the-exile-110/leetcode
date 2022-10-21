package main

/*
pangram 是英文字母中的每个字母至少出现一次的句子。
给定一个仅包含小写英文字母的字符串句子，如果句子是 pangram，则返回 true，否则返回 false。

Example 1:

Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true
Explanation: 句子至少包含英文字母表中的每个字母
Example 2:

Input: sentence = "leetcode"
Output: false
*/

func checkIfPangram(sentence string) bool {
	var m = make(map[rune]bool)
	for _, c := range sentence {
		m[c] = true
	}
	return len(m) == 26
}
