package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9}
	res := findErrorNums(arr)
	fmt.Println("result: ", res)
}

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	duplicated, mismatch := 0, 0

	mp := make(map[int]int)
	for _, v := range nums {
		mp[v] = mp[v] + 1
		if mp[v] == 2 {
			duplicated = v
		}
	}

	for i := 1; i <= len(nums); i++ {
		if mp[i] == 0 {
			mismatch = i
		}
	}

	return []int{duplicated, mismatch}
}
