package main

func main() {

}

func maxDistToClosest(seats []int) int {
	f := make([]int, 0)
	for i, v := range seats {
		if v == 1 {
			f = append(f, i)
		}
	}
	ans := max(f[0]-0, len(seats)-1-f[len(f)-1])
	for i := 1; i < len(f); i++ {
		ans = max(ans, (f[i]-f[i-1])/2)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
