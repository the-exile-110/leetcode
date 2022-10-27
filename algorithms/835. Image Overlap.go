package main

func largestOverlap(A [][]int, B [][]int) int {
	aValues := make([]int, len(A))
	bValues := make([]int, len(B))
	for i, vs := range A {
		for j, v := range vs {
			if v == 1 {
				aValues[i] += 1 << j
			}
			if B[i][j] == 1 {
				bValues[i] += 1 << j
			}
		}
	}
	maxValue := 0
	// i: row offset
	// j: col offset
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			current1, current2 := 0, 0
			for m := 0; m+i < len(A); m++ {
				current1 += getOne(aValues[m] & (bValues[i+m] >> j))
				current2 += getOne(bValues[m] & (aValues[i+m] >> j))
			}
			if current1 > maxValue {
				maxValue = current1
			}
			if current2 > maxValue {
				maxValue = current2
			}
		}
	}
	return maxValue
}

func getOne(v int) int {
	n := 0
	for v != 0 {
		if v&0x1 == 1 {
			n++
		}
		v >>= 1
	}
	return n
}
