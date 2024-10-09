package main

func findSmallestInteger(nums []int, m int) (mex int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[(x%m+m)%m]++
	}
	for cnt[mex%m] > 0 {
		cnt[mex%m]--
		mex++
	}
	return
}
