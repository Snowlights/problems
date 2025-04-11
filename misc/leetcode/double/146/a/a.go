package main

func countSubarrays(nums []int) (ans int) {
	for i := 2; i < len(nums); i++ {
		if (nums[i-2]+nums[i])*2 == nums[i-1] {
			ans++
		}
	}
	return
}
