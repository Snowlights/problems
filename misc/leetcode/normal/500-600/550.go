package _00_600

// 556
func matrixReshape(mat [][]int, r int, c int) [][]int {
	row, col := len(mat), len(mat[0])
	if row*col != r*c {
		return mat
	}
	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	for i := 0; i < row*col; i++ {
		ans[i/c][i%c] = mat[i/col][i%col]
	}
	return ans
}

// 560
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

// 572
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	var bfs func(root, sub *TreeNode) bool
	bfs = func(root, sub *TreeNode) bool {
		if root == nil && sub == nil {
			return true
		}
		if root == nil || sub == nil {
			return false
		}
		if root.Val != sub.Val {
			return false
		}
		return bfs(root.Left, sub.Left) && bfs(root.Right, sub.Right)
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if bfs(v, subRoot) {
				return true
			}
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
	}
	return false
}
