package main

import "fmt"

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)
	var result []int
	for index, num := range nums {
		i, ok := record[target-num]
		if ok {
			return append(result, i, index)
		} else {
			record[num] = index
		}
	}
	return append(result, -1, -1)
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	for _, num := range res {
		fmt.Println(num)
	}
}
