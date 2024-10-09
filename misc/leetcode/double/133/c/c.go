package main

func minOperations(nums []int) (ans int) {
	for _, x := range nums {
		if x == ans%2 {
			ans++
		}
	}
	return
}
