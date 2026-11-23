// Package stock: prices[i] is the stock price on day i. Buy once, then
// sell once on a later day. Return the largest profit, or 0.
//
// Similar to LeetCode 121. Pattern: a single pass that remembers the lowest price so far.
// Time O(n), space O(1).
package stock

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	lowest, profit := prices[0], 0
	for _, p := range prices[1:] {
		profit = max(profit, p-lowest) // max and min are built in since Go 1.21
		lowest = min(lowest, p)
	}
	return profit
}
