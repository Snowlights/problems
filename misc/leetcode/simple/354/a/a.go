package main

func sumOfSquares(nums []int) (ans int) {
	for i, x := range nums {
		if len(nums)%(i+1) == 0 {
			ans += x * x
		}
	}
	return
}
