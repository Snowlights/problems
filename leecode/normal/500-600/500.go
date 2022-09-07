package _00_600

import (
	"container/list"
	"sort"
)

// 504
func nextGreaterElements(nums []int) []int {
	ans, stack := make([]int, len(nums)), make([]int, 0)
	n := len(nums)
	for i := range ans {
		ans[i] = -1
	}
	nums = append(nums, nums...)
	for i, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top%n] = v
		}
		stack = append(stack, i)
	}

	return ans
}

// 509

func fib(n int) int {
	a, b := 0, 1
	switch n {
	case 0:
		return a
	case 1:
		return b
	default:
		for i := 2; i <= n; i++ {
			a, b = b, a+b
		}
		return b
	}
}

// 542
/*
	542. 01 矩阵
	算法思想:BFS(队列——使用go语言的list列表)
*/

func updateMatrix(mat [][]int) [][]int {
	//首先将所有的0的位置入队，并且将1的位置设置成-1，表示表示该位置是未被访问过的1
	queue := list.New()
	m := len(mat)
	n := len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				//入队
				queue.PushBack([]int{i, j})
			} else {
				//表示表示该位置是未被访问过的1
				mat[i][j] = -1
			}
		}
	}
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for queue.Len() != 0 {
		position := queue.Front().Value.([]int) //获取队首元素，并将值强制转换回[]int切片类型
		queue.Remove(queue.Front())
		x := position[0]
		y := position[1]
		for k := 0; k < 4; k++ {
			newX := x + dx[k]
			newY := y + dy[k]
			// 如果四邻域的点是 -1，表示这个点是未被访问过的 1
			// 所以这个点到 0 的距离就可以更新成 matrix[x][y] + 1
			for newX >= 0 && newX < m && newY >= 0 && newY < n && mat[newX][newY] == -1 {
				mat[newX][newY] = mat[x][y] + 1
				queue.PushBack([]int{newX, newY})
			}
		}
	}
	return mat
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 524
func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) ||
			(len(dictionary[i]) == len(dictionary[j]) && dictionary[i] < dictionary[j])
	})

	for _, v := range dictionary {
		sbegin, vbegin := 0, 0
		for sbegin < len(s) && vbegin < len(v) {
			if s[sbegin] == v[vbegin] {
				sbegin++
				vbegin++
				continue
			}
			sbegin++
		}
		if vbegin == len(v) {
			return v
		}
	}
	return ""
}
