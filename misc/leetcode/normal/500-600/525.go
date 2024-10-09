package _00_600

import (
	"container/list"
	"fmt"
	"math/rand"
	"sort"
)

// 528
type Solution struct {
	prefix []int
}

func Constructor(w []int) Solution {
	prefix := make([]int, len(w))
	prefix[0] = w[0]
	for i := 1; i < len(w); i++ {
		prefix[i] = prefix[i-1] + w[i]
	}
	return Solution{prefix}
}

func (this *Solution) PickIndex() int {
	val := rand.Intn(this.prefix[len(this.prefix)-1]) + 1
	idx := sort.SearchInts(this.prefix, val)
	fmt.Println(val, idx)
	return this.prefix[idx]
}

// 529
func updateBoard(board [][]byte, click []int) [][]byte {
	m, n := len(board), len(board[0])
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	dir := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	q := [][]int{{x, y}}
	vis[x][y] = true
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			boom := '0'
			for _, d := range dir {
				x, y := v[0]+d[0], v[1]+d[1]
				if 0 <= x && x < m && 0 <= y && y < n {
					if board[x][y] == 'M' {
						boom++
					}
				}
			}
			if boom == '0' {
				board[v[0]][v[1]] = 'B'
				for _, d := range dir {
					x, y := v[0]+d[0], v[1]+d[1]
					if 0 <= x && x < m && 0 <= y && y < n {
						if board[x][y] == 'e' && !vis[x][y] {
							q = append(q, []int{x, y})
							vis[x][y] = true
						}
					}
				}
			} else {
				board[v[0]][v[1]] = byte(boom)
			}
		}
	}

	return board
}

// 540
func singleNonDuplicate(a []int) int {
	i := sort.Search(len(a)-1, func(i int) bool {
		return a[i] != a[i^1]
	})
	return a[i]
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

// 543
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	var dfs func(node *TreeNode) int
	ans := 0
	dfs = func(from *TreeNode) int {
		if from == nil {
			return 0
		}
		if from.Left == nil && from.Right == nil {
			return 1
		}
		l, r := dfs(from.Left), dfs(from.Right)
		ans = max(ans, l+r)
		return max(l, r) + 1
	}

	dfs(root)
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 547
func findCircleNum(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !vis[to] {
				dfs(to)
			}
		}
	}
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return
}
