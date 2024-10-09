package _00_400

import (
	"math"
	"sort"
)

// 367
func isPerfectSquare(num int) bool {
	x := int(math.Sqrt(float64(num)))
	return x*x == num
}

// 374
func guessNumber(n int) int {
	var guess func(num int) int
	return sort.Search(n, func(x int) bool { return guess(x) <= 0 })
}
