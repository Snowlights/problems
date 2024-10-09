package main

func numberOfAlternatingGroups2(colors []int) int {
	ans, n := 0, len(colors)
	for i, v := range colors {
		if colors[(i+1)%n] != v && colors[(i+1)%n] == colors[(i-1+n)%n] {
			ans++
		}
	}
	return ans
}
