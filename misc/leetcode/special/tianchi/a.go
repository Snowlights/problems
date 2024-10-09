package tianchi

// https://leetcode.cn/contest/tianchi2022/

// 1
type ListNode struct {
	Val  int
	Next *ListNode
}

func numberEvenListNode(head *ListNode) int {
	ans := 0
	for head != nil {
		if head.Val%2 == 1 {
			ans++
		}
		head = head.Next
	}
	return ans
}

// 2
func getLength(grid []string) int {
	m, n := len(grid), len(grid[0])
	const (
		up    = 1
		down  = 2
		left  = 3
		right = 4
	)
	x, y, d := 0, 0, down
	ans := 0
	for 0 <= x && x < m && 0 <= y && y < n {
		ans++
		switch grid[x][y] {
		case 'L':
			switch d {
			case up:
				d = left
				y--
			case down:
				d = right
				y++
			case left:
				d = up
				x--
			case right:
				d = down
				x++
			}
		case 'R':
			switch d {
			case up:
				d = right
				y++
			case down:
				d = left
				y--
			case left:
				d = down
				x++
			case right:
				d = up
				x--
			}
		default:
			switch d {
			case up:
				x--
			case down:
				x++
			case left:
				y--
			case right:
				y++
			}
		}
	}

	return ans
}

// 3
func arrangeBookshelf(order []int, limit int) []int {
	c, h := make(map[int]int), make(map[int]int)
	for _, v := range order {
		c[v]++
	}
	st := []int{0}
	for _, v := range order {
		if h[v] == limit {
			c[v]--
			continue
		}
		for v < st[len(st)-1] && c[st[len(st)-1]] > limit {
			c[st[len(st)-1]]--
			h[st[len(st)-1]]--
			st = st[:len(st)-1]
		}
		st = append(st, v)
		h[v]++
	}

	return st[1:]
}

// 4
func brilliantSurprise(a [][]int, lim int) (ans int) {
	dp := make([]int, lim+1)
	var f func([][]int, []int)
	f = func(a [][]int, tot []int) {
		if len(a) == 1 {
			s := 0
			for i, v := range a[0] {
				if i >= lim {
					break
				}
				s += v
				ans = max(ans, dp[lim-(i+1)]+s)
			}
			return
		}

		tmp := append([]int{}, dp...)

		m := len(a) / 2
		for i, r := range a[:m] {
			for j := lim; j >= len(r); j-- {
				dp[j] = max(dp[j], dp[j-len(r)]+tot[i])
			}
		}
		f(a[m:], tot[m:])

		dp = tmp
		for i, r := range a[m:] {
			for j := lim; j >= len(r); j-- {
				dp[j] = max(dp[j], dp[j-len(r)]+tot[m+i])
			}
		}
		f(a[:m], tot[:m])
	}

	tot := make([]int, len(a))
	for i, r := range a {
		for _, v := range r {
			tot[i] += v
		}
	}
	f(a, tot)
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
