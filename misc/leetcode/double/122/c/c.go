package main

func minimumArrayLength(nums []int) int {
	mm, cnt := nums[0], 0
	for _, v := range nums {
		if v < mm {
			mm = v
		}
	}
	for _, v := range nums {
		if v%mm > 0 {
			return 1
		}
		if v == mm {
			cnt++
		}
	}
	return (cnt + 1) / 2
}
