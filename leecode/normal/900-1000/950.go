package _00_1000

import "sort"

// 968
var res int

func minCameraCover(root *TreeNode) int {

	res = 0
	if dfs(root) == 1 {
		// 如果最顶端的元素没有被安装，也没有被监视
		// 那么这个节点需要安装一个摄像头
		res++
	}
	return res
}

// 1未被监视, 2安装摄像头
func dfs(root *TreeNode) int {

	// 最后一个节点不关心
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	// 如果子节点没有被监视
	if left == 1 || right == 1 {
		// 安装摄像头
		res++
		return 2
	}

	// 如果子节点安装摄像头
	if left == 2 || right == 2 {
		// 标记被监视
		return 3
	}
	return 1
}

// 973
func kClosest(points [][]int, k int) [][]int {

	type node struct {
		x, y, num int
	}
	nodeList := make([]*node, 0, len(points))
	for _, p := range points {
		nodeList = append(nodeList, &node{
			x:   p[0],
			y:   p[1],
			num: p[0]*p[0] + p[1]*p[1],
		})
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].num < nodeList[j].num
	})
	ans := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, []int{nodeList[i].x, nodeList[i].y})
	}
	return ans
}
