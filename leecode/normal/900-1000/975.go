package _00_1000

import "sort"

// 977
func sortedSquares(nums []int) []int {

	ans := make([]int, 0, len(nums))
	for _, v := range nums {
		ans = append(ans, v*v)
	}
	sort.Ints(ans)
	return ans
}

// 981
type TimeMap struct {
	h map[string][]*pair
}

type pair struct {
	val       string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{h: make(map[string][]*pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	this.h[key] = append(this.h[key], &pair{
		val:       value,
		timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	idx := sort.Search(len(this.h[key]), func(i int) bool {
		return this.h[key][i].timestamp > timestamp
	})
	if 0 < idx && idx <= len(this.h[key]) {
		return this.h[key][idx-1].val
	}
	return ""
}

// 986
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int, 0)

	for _, v := range firstList {
		idx := sort.Search(len(secondList), func(i int) bool {
			return secondList[i][0] >= v[0]
		})
		if idx == len(secondList) {
			continue
		}
		base := []int{v[0], v[1]}
		for idx < len(secondList) && secondList[idx][0] >= base[0] {
			base[0] = max(base[0], secondList[idx][0])
			base[1] = min(base[1], secondList[idx][1])
			idx++
		}
		ans = append(ans, base)
	}

	return ans
}

// 989
func addToArrayForm(num []int, k int) (ans []int) {
	for i := len(num) - 1; i >= 0; i-- {
		sum := num[i] + k%10
		k /= 10
		if sum >= 10 {
			k++
			sum -= 10
		}
		ans = append(ans, sum)
	}
	for ; k > 0; k /= 10 {
		ans = append(ans, k%10)
	}
	reserve(ans)
	return
}

func reserve(num []int) []int {
	s, e := 0, len(num)-1
	for s < e {
		num[s], num[e] = num[e], num[s]
		s++
		e--
	}
	return num
}

// 994
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	q, ans := make([][]int, 0), 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			x, y := v[0], v[1]
			for _, d := range dir {
				xx, yy := x+d[0], y+d[1]
				if 0 <= xx && xx < m && 0 <= yy && yy < n && grid[xx][yy] == 1 {
					grid[xx][yy] = 2
					q = append(q, []int{xx, yy})
				}
			}
		}
		if len(q) > 0 {
			ans++
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans
}

// 998
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	for cur := root; cur != nil; cur = cur.Right {
		if val > cur.Val {
			if parent == nil {
				return &TreeNode{val, root, nil}
			}
			parent.Right = &TreeNode{val, cur, nil}
			return root
		}
		parent = cur
	}
	parent.Right = &TreeNode{Val: val}
	return root
}
