package double_102

import (
	"math"
	"strconv"
)

// 1
func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := range grid[0] {
		for _, row := range grid {
			ans[j] = max(ans[j], len(strconv.Itoa(row[j])))
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 2
func findPrefixScore(nums []int) []int64 {
	ans := make([]int64, len(nums))
	mx, s := 0, 0
	for i, x := range nums {
		mx = max(mx, x) // 前缀最大值
		s += x + mx     // 累加前缀的得分
		ans[i] = int64(s)
	}
	return ans
}

// 3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	root.Val = 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		nextLevelSum := 0 // 下一层的节点值之和
		for _, node := range tmp {
			if node.Left != nil {
				q = append(q, node.Left)
				nextLevelSum += node.Left.Val
			}
			if node.Right != nil {
				q = append(q, node.Right)
				nextLevelSum += node.Right.Val
			}
		}

		// 再次遍历，更新下一层的节点值
		for _, node := range tmp {
			childrenSum := 0 // node 左右儿子的节点值之和
			if node.Left != nil {
				childrenSum += node.Left.Val
			}
			if node.Right != nil {
				childrenSum += node.Right.Val
			}
			// 更新 node 左右儿子的节点值
			if node.Left != nil {
				node.Left.Val = nextLevelSum - childrenSum
			}
			if node.Right != nil {
				node.Right.Val = nextLevelSum - childrenSum
			}
		}
	}
	return root
}

// 4
const inf = math.MaxInt32 / 2 // 防止更新最短路时加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n) // 邻接矩阵
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边）
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]][e[1]] = e[2] // 添加一条边（输入保证这条边之前不存在）
}

// 朴素 Dijkstra 算法
func (g Graph) ShortestPath(start, end int) int {
	n := len(g)
	dis := make([]int, n) // 从 start 出发，到各个点的最短路，如果不存在则为无穷大
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	vis := make([]bool, n)
	for {
		// 找到当前最短路，去更新它的邻居的最短路，
		// 根据数学归纳法，dis[x] 一定是最短路长度
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 || dis[x] == inf { // 所有从 start 能到达的点都被更新了
			return -1
		}
		if x == end { // 找到终点，提前退出
			return dis[x]
		}
		vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range g[x] {
			dis[y] = min(dis[y], dis[x]+w) // 更新最短路长度
		}
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// Floyd
//const inf = math.MaxInt / 3 // 防止更新最短路时加法溢出
//
//type Graph [][]int
//
//func Constructor(n int, edges [][]int) Graph {
//	d := make([][]int, n) // 邻接矩阵
//	for i := range d {
//		d[i] = make([]int, n)
//		for j := range d[i] {
//			if j != i {
//				d[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
//			}
//		}
//	}
//	for _, e := range edges {
//		d[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边和自环）
//	}
//	for k := range d {
//		for i := range d {
//			for j := range d {
//				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
//			}
//		}
//	}
//	return d
//}
//
//func (d Graph) AddEdge(e []int) {
//	x, y, w := e[0], e[1], e[2]
//	if w >= d[x][y] { // 无需更新
//		return
//	}
//	for i := range d {
//		for j := range d {
//			d[i][j] = min(d[i][j], d[i][x]+w+d[y][j])
//		}
//	}
//}
//
//func (d Graph) ShortestPath(start, end int) int {
//	ans := d[start][end]
//	if ans == inf {
//		return -1
//	}
//	return ans
//}
