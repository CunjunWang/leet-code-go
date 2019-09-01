package main

import "fmt"

func maxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func twoSumLessThanK(A []int, K int) int {
	size := len(A)
	maxLessThanK := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			sum := A[i] + A[j]
			if sum < K {
				maxLessThanK = maxInt(maxLessThanK, sum)
			}
		}
	}
	if maxLessThanK == 0 {
		return -1
	} else {
		return maxLessThanK
	}
}

func main() {
	A := []int{10, 20, 30}
	K := 15
	res := twoSumLessThanK(A, K)
	fmt.Println(res)
}
