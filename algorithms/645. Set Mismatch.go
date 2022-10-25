package main

func findErrorNums(nums []int) []int {
	result := make([]int, 2)
	size := len(nums)
	sum := 0
	existNums := make([]bool, size+1)
	for _, num := range nums {
		sum += num
		if !existNums[num] {
			existNums[num] = true
			continue
		}
		result[0] = num
	}

	result[1] = size*(size+1)/2 - sum + result[0]

	return result
}
