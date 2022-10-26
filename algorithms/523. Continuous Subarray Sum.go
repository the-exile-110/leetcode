package main

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	if k == 0 {
		for i := 1; i < len(nums); i++ {
			if nums[i] == 0 && nums[i-1] == 0 {
				return true
			}
		}
		return false
	}
	if k < 0 {
		k = -k
	}
	var sum int
	m := make(map[int]int)
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sum %= k
		if j, ok := m[sum]; ok {
			if i-j > 1 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}
