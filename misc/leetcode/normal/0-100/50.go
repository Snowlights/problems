package __100

import (
	"sort"
	"strconv"
	"strings"
)

// 51
func solveNQueens(n int) (ans [][]string) {
	col := make([]int, n)
	onPath := make([]bool, n)
	diag1 := make([]bool, n*2-1)
	diag2 := make([]bool, n*2-1)
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range col {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			ans = append(ans, board)
			return
		}
		for c, on := range onPath {
			rc := r - c + n - 1
			if !on && !diag1[r+c] && !diag2[rc] {
				col[r] = c
				onPath[c], diag1[r+c], diag2[rc] = true, true, true
				dfs(r + 1)
				onPath[c], diag1[r+c], diag2[rc] = false, false, false // 恢复现场
			}
		}
	}
	dfs(0)
	return
}

// 53
func maxSubArray(nums []int) int {
	ans, tmp := nums[0], 0
	for _, v := range nums {
		tmp += v
		if tmp > ans {
			ans = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	return ans
}

// 54
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)

	xS, xE, yS, yE := 0, m-1, 0, n-1
	for xS <= xE && yS <= yE {
		s, e := yS, yE
		for i := s; i <= e; i++ {
			ans = append(ans, matrix[xS][i])
		}
		s, e = xS+1, xE
		for i := s; i <= e; i++ {
			ans = append(ans, matrix[i][yE])
		}
		if xS < xE && yS < yE {
			s, e = yE-1, yS
			for i := s; i >= e; i-- {
				ans = append(ans, matrix[xE][i])
			}
			s, e = xE-1, xS+1
			for i := s; i >= e; i-- {
				ans = append(ans, matrix[i][xS])
			}
		}
		xS++
		xE--
		yS++
		yE--
	}

	return ans
}

// 56
func merge56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	s, e := 0, len(intervals)
	ans := make([][]int, 0)
	for s < e {
		node := []int{intervals[s][0], intervals[s][1]}
		s++
		for s < e {
			if node[0] <= intervals[s][0] && intervals[s][0] <= node[1] {
				node[0] = min(intervals[s][0], node[0])
				node[1] = max(intervals[s][1], node[1])
			} else {
				break
			}
			s++
		}
		ans = append(ans, node)
	}
	return ans
}

// 57
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i][1] = max(intervals[i][1], intervals[i+1][1]) //赋值最大值
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}
	}
	return intervals
}

// 61
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	ll := 0
	cur := head
	for cur != nil {
		ll++
		cur = cur.Next
	}
	k %= ll
	if k == 0 {
		return head
	}
	cur = head
	k = ll - k - 1
	for k > 0 {
		cur = cur.Next
		k--
	}
	next := cur.Next
	cur.Next = nil
	cur = next
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return next
}

// 62
func uniquePaths(m int, n int) int {

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 63
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			continue
		}
		dp[0][i] = dp[0][i-1]
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			continue
		}
		dp[i][0] = dp[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 64
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// 67
func addBinary(a string, b string) string {
	res := ""
	carry := 0
	l1, l2 := len(a)-1, len(b)-1
	for l1 >= 0 || l2 >= 0 {
		x, y := 0, 0
		if l1 >= 0 {
			x = int(a[l1] - byte('0'))
		}
		if l2 >= 0 {
			y = int(b[l2] - byte('0'))
		}

		sum := x + y + carry
		res += strconv.Itoa(sum % 2)
		carry = sum / 2

		l1--
		l2--
	}
	if carry != 0 {
		res += strconv.Itoa(carry)
	}

	return reverseString(res)
}

func reverseString(str string) string {
	temp := []byte(str)
	left, right := 0, len(temp)-1
	for left < right {
		temp[left], temp[right] = temp[right], temp[left]
		left++
		right--
	}
	return string(temp)
}

// 70
func climbStairs(n int) int {
	f1, f2 := 1, 2
	if n == 1 {
		return f1
	}
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}

// 72
func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1) //dp数组存放上一次dp的情况
	for i := 1; i < len(word2)+1; i++ {
		dp[i] = i //①:对第1行进行初始化,此时取word1前0个字母与word2前i个字母的编辑距离 即word2的长度
	}
	for i := 1; i < len(word1)+1; i++ {
		leftUp := i - 1 //用于存储dp[i][j]左上点的数据(因为会被dp[i-1][j]覆盖掉)
		dp[0] = i       //初始化每一行的首位元素 类似①的情况
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] { //当此处字符相同 则通过替换的转移方式可以实现无替换转移,即不变
				dp[j], leftUp = minV2(dp[j]+1, dp[j-1]+1, leftUp), dp[j]
			} else {
				dp[j], leftUp = minV2(dp[j]+1, dp[j-1]+1, leftUp+1), dp[j]
			}
		}
	}
	return dp[len(word2)]
}

func minV2(a, b, c int) (ret int) {
	ret = a
	if a > b {
		ret = b
		if b > c {
			ret = c
		}
	} else if a > c {
		ret = c
	}
	return
}

// 74
func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if row[0] <= target && target <= row[len(row)-1] {
			idx := sort.Search(len(row), func(i int) bool {
				return row[i] >= target
			})
			return row[idx] == target
		}
	}
	return false
}

// 77
func combine(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}
