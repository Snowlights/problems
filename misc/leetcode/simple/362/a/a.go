package main

func numberOfPoints(nums [][]int) (ans int) {
	diff := [102]int{}
	for _, p := range nums {
		diff[p[0]]++
		diff[p[1]+1]--
	}
	s := 0
	for _, d := range diff {
		s += d
		if s > 0 {
			ans++
		}
	}
	return
}
