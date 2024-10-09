package main

func countTestedDevices(b []int) int {
	ans := 0
	for _, v := range b {
		if v > ans {
			ans++
		}
	}
	return ans
}
