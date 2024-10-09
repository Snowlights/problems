package _600_1700

import "math"

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

// 1620
func bestCoordinate(towers [][]int, radius int) []int {
	var xMax, yMax, cx, cy, maxQuality int
	for _, t := range towers {
		xMax = max(xMax, t[0])
		yMax = max(yMax, t[1])
	}
	for x := 0; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			quality := 0
			for _, t := range towers {
				d := (x-t[0])*(x-t[0]) + (y-t[1])*(y-t[1])
				if d <= radius*radius {
					quality += int(float64(t[2]) / (1 + math.Sqrt(float64(d))))
				}
			}
			if quality > maxQuality {
				cx, cy, maxQuality = x, y, quality
			}
		}
	}
	return []int{cx, cy}
}
