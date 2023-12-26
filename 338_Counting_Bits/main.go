package main

import "fmt"

func main() {
	result := countBits(2)
	fmt.Printf("result: %+v\n", result)
}

// O(n log n)
// run time:        3 ms
// Memory usage: 6.36 MB bcuz of string conversion
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
// run time:        0 ms lol
// Memory usage: 4.46 MB
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
// run time:        6 ms
// Memory usage: 4.48 MB
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

// O (n log n)
// but reduce memory usage by removing the string conversion
// run time:        7 ms
// Memory usage: 4.48 MB
func countBits3(n int) []int {
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
