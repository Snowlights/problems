package _000_1100

// 1042

func gardenNoAdj(n int, paths [][]int) []int {
	var res = make([]int, n)
	// 构建邻接矩阵
	var c = make([][]int, n)
	for _, line := range paths {
		c[line[0]-1] = append(c[line[0]-1], line[1]-1)
		c[line[1]-1] = append(c[line[1]-1], line[0]-1)
	}
	for i := 0; i < n; i++ {
		used := make(map[int]struct{})
		for _, point := range c[i] {
			if res[point] != 0 {
				used[res[point]] = struct{}{}
			}
		}
		for colour := 1; colour <= 4; colour++ {
			if _, ok := used[colour]; !ok {
				res[i] = colour
				break
			}
		}
	}
	return res
}
