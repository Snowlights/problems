package _00_300

import "strconv"

// 226
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 227
func calculate227(s string) (ans int) {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}

// 230
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

// 233
func countDigitOne(n int) int {
	s := strconv.Itoa(n)

	dp := make([][]int64, len(s))
	for i := range dp {
		dp[i] = make([]int64, len(s))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var f func(p, sum int, limitUp bool) int64
	f = func(p, sum int, limitUp bool) (res int64) {
		if p == len(s) {
			return int64(sum)
		} // sum
		if !limitUp {
			dv := &dp[p][sum]
			if *dv >= 0 {
				return *dv
			} // *dv + sum*int64(math.Pow10(n-p))
			defer func() { *dv = res }()
		}
		up := 9
		if limitUp {
			up = int(s[p] - '0')
		}
		for ch := 0; ch <= up; ch++ {
			c := sum
			if ch == 1 {
				c++
			}
			res += f(p+1, c, limitUp && ch == up)
		}
		return
	}
	return int(f(0, 0, true))
}

// 235
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	stoP, valToNode := make(map[int]*TreeNode), make(map[int]*TreeNode)
	que := []*TreeNode{root}
	stoP[root.Val] = root
	for len(que) > 0 {
		tmp := que
		que = nil
		for _, v := range tmp {
			valToNode[v.Val] = v
			if v.Left != nil {
				stoP[v.Left.Val] = v
				que = append(que, v.Left)
			}
			if v.Right != nil {
				stoP[v.Right.Val] = v
				que = append(que, v.Right)
			}
		}
	}

	pn, qn := valToNode[p.Val], valToNode[q.Val]
	pMap := make(map[int]bool)
	pMap[pn.Val] = true
	for {
		parent := stoP[pn.Val]
		pMap[parent.Val] = true
		if parent == pn {
			break
		}
		pn = parent
	}
	if pMap[qn.Val] {
		return valToNode[qn.Val]
	}
	for {
		parent := stoP[qn.Val]
		if pMap[parent.Val] {
			return valToNode[parent.Val]
		}
		qn = parent
	}

}

// 239
func maxSlidingWindow(nums []int, k int) []int {
	q, ans := []int{}, []int{}

	for i, v := range nums {
		for len(q) > 0 && v >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		for len(q) > 0 && q[0] <= i-k {
			q = q[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}

	return ans
}
