package main

import "fmt"

func main() {
	result := countBits(2)
	fmt.Printf("result: %+v\n", result)
}

// O(n log n)
func countBits(n int) []int {
	bitArr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bin := fmt.Sprintf("%b\n", i)
		count := 0
		for _, char := range bin {
			if char == '1' {
				count++
			}
		}
		bitArr[i] = count
	}
	return bitArr
}

// O(n)
// suggest by Github Copilot
func countBitsOn(n int) []int {
	bitArr := make([]int, n+1)
	bitArr[0] = 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			bitArr[i] = bitArr[i/2]
		} else {
			bitArr[i] = bitArr[i-1] + 1
		}
	}
	return bitArr
}

// O(nm)
// at first github copilot say my first attempt is O(nm) but I don't know what is, so i let it show me this method
func countBitsNM(n int) []int {
	bitArr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		for j := i; j > 0; j = j / 2 {
			if j%2 == 1 {
				count++
			}
		}
		bitArr[i] = count
	}
	return bitArr
}
