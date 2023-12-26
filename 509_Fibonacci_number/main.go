package main

func main() {
	fib(6)
}

// what is fib lol
// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.

// i'm studying dynamic programming and this way has time complexity of O(n)
// run time:        1 ms
// Memory usage: 1.95 MB
func fib(n int) int {
	fibs := make([]int, n+1)
	fibs[0], fibs[1] = 0, 1

	for i := 2; i <= n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs[n]
}

// how to write in a recursive way (O(2^n))
// run time:        9 ms
// Memory usage: 1.95 MB
func fibRec(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// in iterative way (O(n)) (Copilot suggested this so i wanna remember this)
// run time:        1 ms
// Memory usage: 1.96 MB
func fibIter(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Note: Diff between dynamic programming and iterative way is
// dynamic programming uses memoization to store the result of subproblems
// and use it later on, while iterative way doesn't store the result of subproblems
// and just calculate it again and again.
