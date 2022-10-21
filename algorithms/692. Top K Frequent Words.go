package main

import "sort"

func topKFrequent(words []string, k int) []string {
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	wordList := make([]string, 0, len(wordCount))
	for word := range wordCount {
		wordList = append(wordList, word)
	}
	sort.Slice(wordList, func(i, j int) bool {
		if wordCount[wordList[i]] == wordCount[wordList[j]] {
			return wordList[i] < wordList[j]
		}
		return wordCount[wordList[i]] > wordCount[wordList[j]]
	})
	return wordList[:k]
}
