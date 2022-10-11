package _600_1700

// 1608
func specialArray(nums []int) int {

	for i := 1; i <= 1005; i++ {
		tmp := 0
		for _, v := range nums {
			if v >= i {
				tmp++
			}
		}

		if tmp == i {
			return i
		}
	}
	return -1
}

// 1615
func maximalNetworkRank(n int, roads [][]int) int {
	m := make([][]bool, n)
	for i := range m {
		m[i] = make([]bool, n)
	}
	indegree := make(map[int]int)

	for _, r := range roads {
		a, b := r[0], r[1]
		indegree[a]++
		indegree[b]++
		m[a][b] = true
		m[b][a] = true
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := indegree[i] + indegree[j]
			if m[i][j] || m[j][i] {
				tmp--
			}
			ans = max(ans, tmp)
		}
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
