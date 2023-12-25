package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	decodeCnt := make([]int, n+1)
	decodeCnt[0] = 1 // Base case: empty string has one way to decode

	for i := 1; i <= n; i++ {
		// every loop can be decoded as a single digit and a double digit, that why no need  to use if else
		if s[i-1] != '0' {
			decodeCnt[i] += decodeCnt[i-1]
			fmt.Println("for debug")
		}
		if i > 1 {
			twoDigits, err := strconv.Atoi(s[i-2 : i])
			if twoDigits >= 10 && twoDigits <= 26 && err == nil {
				decodeCnt[i] += decodeCnt[i-2]
				fmt.Println("for debug")
			}
		}
	}

	return decodeCnt[n]
}

func main() {
	fmt.Println("result: ", numDecodings("226"))
}

// great article for understand Dynamic programing https://medium.com/@aquablitz11/an-introduction-to-dynamic-programming-76574fda6501
