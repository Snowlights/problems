package main

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i, x := range nums {
		result[i] = nums[((i+x)%n+n)%n] // 保证结果在 [0,n-1] 中
	}
	return result
}
