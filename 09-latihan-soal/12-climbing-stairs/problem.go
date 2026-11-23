// Package climbingstairs: a staircase has n steps. Each move climbs 1 or 2 steps.
// How many distinct ways are there to reach the top?
//
// Similar to LeetCode 70. Pattern: ways(n) = ways(n-1) + ways(n-2), like Fibonacci.
package climbingstairs

// climbRecursive follows the formula directly. Correct, but time O(2^n)
// because the same values are computed again and again.
func climbRecursive(n int) int {
	if n <= 2 {
		return n
	}
	return climbRecursive(n-1) + climbRecursive(n-2)
}

// climbMemo stores results that were already computed. Time O(n), space O(n).
func climbMemo(n int) int {
	memo := map[int]int{}
	var f func(int) int
	f = func(k int) int {
		if k <= 2 {
			return k
		}
		if v, ok := memo[k]; ok {
			return v
		}
		memo[k] = f(k-1) + f(k-2)
		return memo[k]
	}
	return f(n)
}

// climbStairs only remembers the last two values. Time O(n), space O(1).
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
