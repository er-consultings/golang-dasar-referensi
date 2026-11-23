// Package discount is a small example for learning table-driven tests.
//
//	go test -v ./07-test-dan-konkurensi/discount
package discount

// Apply returns the total after taking off the given percentage.
// Zero or negative percentages are ignored, and percentages above 100 are capped at 100.
func Apply(total, percent int64) int64 {
	if percent <= 0 {
		return total
	}
	if percent > 100 {
		percent = 100
	}
	return total - total*percent/100
}
