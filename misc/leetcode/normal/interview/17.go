package interview

import (
	"sort"
)

// 17. 08
// 有个马戏团正在设计叠罗汉的表演节目，一个人要站在另一人的肩膀上。
// 出于实际和美观的考虑，在上面的人要比下面的人矮一点且轻一点。
// 已知马戏团每个人的身高和体重，请编写代码计算叠罗汉最多能叠几个人。

// height = [65,70,56,75,60,68] weight = [100,150,90,190,95,110]
func bestSeqAtIndex(height []int, weight []int) int {
	type pair struct {
		w, h int
	}
	pList := make([]*pair, 0)
	for i := range height {
		pList = append(pList, &pair{
			w: weight[i],
			h: height[i],
		})
	}

	sort.Slice(pList, func(i, j int) bool {
		return pList[i].h > pList[j].h ||
			(pList[i].h == pList[j].h && pList[i].w < pList[j].w)
	})
	s := []*pair{}

	for _, v := range pList {
		j := sort.Search(len(s), func(i int) bool {
			c := s[i]
			return c.h <= v.h || c.w <= v.w
		})
		// 将第j个人替换成p
		if j == len(s) {
			s = append(s, v)
		} else {
			s[j] = v
		}
	}
	return len(s)
}

// 17. 23
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findSquare(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}

	// 预先计算，对于每一行，当以某个位置i结尾时，最大的连续黑色的长度。
	dp1 := make([][]int, m)
	for i := 0; i < m; i++ {
		dp1[i] = make([]int, n)

		prev := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				prev++
				dp1[i][j] = prev
			} else {
				prev = 0
			}
		}
	}

	//  预先计算，对于每一列，当以某个位置j结尾时，最大的连续黑色的长度
	dp2 := make([][]int, n)
	for j := 0; j < n; j++ {
		dp2[j] = make([]int, m)

		prev := 0
		for i := 0; i < m; i++ {
			if matrix[i][j] == 0 {
				prev++
				dp2[j][i] = prev
			} else {
				prev = 0
			}
		}
	}

	// 找最大矩形
	var maxSize int = 0
	var x, y int // 右下角的点
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//i和j是右下角，尝试每一个可能的长度，最大的可能长度是两个边长的最小值
			size := min(dp1[i][j], dp2[j][i])
			for size > 0 {
				// 查看 矩形top那一行和left那一列是不是也都是连续黑色
				if dp1[i-size+1][j] >= size && dp2[j-size+1][i] >= size {
					break
				}
				size--
			}

			if size > maxSize {
				x, y = i, j
				maxSize = size
			}

		}
	}

	if maxSize == 0 { // 全白就返回空
		return []int{}
	}

	return []int{x - maxSize + 1, y - maxSize + 1, maxSize} // 要返回左上角
}
