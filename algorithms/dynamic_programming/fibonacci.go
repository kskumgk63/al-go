package dynamic_programming

func fib(n int, memo map[int]int) int {
	if n < 2 {
		return n
	}
	if _, ok := memo[n]; !ok {
		return fib(n-1, memo) + fib(n-2, memo)
	}
	return memo[n]
}
