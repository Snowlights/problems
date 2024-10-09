package offer

import (
	"math"
	"strings"
)

// offer 04
func findNumberIn2DArray(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	for i, j := 0, len(matrix[0])-1; i < row && j >= 0; {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// offer 05
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// offer 06
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	s, e := 0, len(arr)-1
	for s < e {
		arr[s], arr[e] = arr[e], arr[s]
		s++
		e--
	}
	return arr
}

// offer 09
type CQueue struct {
	list []int
}

func Constructor9() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.list = append(this.list, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.list) == 0 {
		return -1
	}
	val := this.list[0]
	this.list = this.list[1:]
	return val
}

// offer 10
func fib(n int) int {
	a, b := 0, 1
	if n == 1 {
		return a
	}
	mod := int(1e9 + 7)
	for i := 2; i <= n; i++ {
		a, b = b%mod, (a+b)%mod
	}
	return b
}

// offer 12
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

	var bfs func(i, j, idx int) bool
	bfs = func(i, j, idx int) bool {
		if board[i][j] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] {
				vis[x][y] = true
				if bfs(x, y, idx+1) {
					return true
				}
				vis[x][y] = false
			}
		}
		return false
	}

	for i := range board {
		for j := range board[i] {
			vis[i][j] = true
			if bfs(i, j, 0) {
				return true
			}
			vis[i][j] = false
		}
	}
	return false
}

// offer 14-1
func cuttingRope(n int) int {
	// dp[i]: 标识长度为 i 的绳子剪断后能获取的最大乘积,dp[n]即为所求
	dp := make([]int, n+1)

	// 初始化
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		var maxLength int
		for j := 1; j < i; j++ {
			maxLength = max(maxLength, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = maxLength
	}

	return dp[n]
}

// offer 14-2
func cuttingRope2(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	ans := 1
	const mod = 1e9 + 7
	threeTimes, one, two := 0, false, false
	if n%3 == 0 {
		threeTimes = n / 3
	} else if n%3 == 1 {
		threeTimes = n/3 - 1
		one = true
	} else {
		threeTimes = n / 3
		two = true
	}

	for threeTimes > 0 {
		ans = ans * 3 % mod
		threeTimes--
	}
	if one {
		ans = ans * 4 % mod
	}
	if two {
		ans = ans * 2 % mod
	}

	return ans
}

// offer 16
func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

// offer 18
func deleteNode(head *ListNode, val int) *ListNode {

	if head.Val == val {
		return head.Next
	}
	pre, cur := head, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur.Next = nil
			break
		}
		pre = cur
		cur = cur.Next
	}

	return head
}

// offer 21
func exchange(nums []int) []int {
	s, e := 0, len(nums)-1
	for s < e {
		for s < e && nums[s]%2 == 1 {
			s++
		}
		for s < e && nums[e]%2 == 0 {
			e--
		}
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
	return nums
}

// offer 22
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// offer 24
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
