package _00_800

import (
	"sort"
)

// 700
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	for root != nil {
		if root.Val == val {
			return root
		}

		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return nil
}

// 704
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// 706
type MyHashMap struct {
	h map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{make(map[int]int)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.h[key] = value
}

func (this *MyHashMap) Get(key int) int {
	val, ok := this.h[key]
	if ok {
		return val
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.h, key)
}

// 713
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	prod, i := 1, 0
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += j - i + 1
	}
	return
}

// 718
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	ans := 0
	for i, v := range nums1 {
		for j, vv := range nums2 {
			if v == vv {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = 0
			}
			ans = int(max(int64(ans), int64(dp[i+1][j+1])))
		}
	}
	// fmt.Println(dp)
	return ans
}

// 720
func longestWord(words []string) (longest string) {
	sort.Slice(words, func(i, j int) bool {
		s, t := words[i], words[j]
		return len(s) < len(t) || len(s) == len(t) && s > t
	})

	candidates := map[string]struct{}{"": {}}
	for _, word := range words {
		if _, ok := candidates[word[:len(word)-1]]; ok {
			longest = word
			candidates[word] = struct{}{}
		}
	}
	return longest
}

// 721
func accountsMerge(accounts [][]string) [][]string {

	h := make(map[string]int)
	fa := make([]int, len(accounts))
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(from int) int {
		if fa[from] != from {
			fa[from] = find(fa[from])
		}
		return fa[from]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	for i, v := range accounts {
		for _, a := range v[1:] {
			vv, ok := h[a]
			if ok {
				merge(i, vv)
			}
			h[a] = i
		}
	}

	ans := make(map[int]map[string]bool)
	for i, v := range accounts {
		to := find(i)
		m, ok := ans[to]
		if !ok {
			m = make(map[string]bool)
			ans[to] = m
		}
		for _, v := range v[1:] {
			m[v] = true
		}
	}
	res := make([][]string, 0, len(ans))
	for i, v := range ans {
		e := make([]string, 0, len(v))
		for vv := range v {
			e = append(e, vv)
		}
		sort.Strings(e)
		e = append([]string{accounts[i][0]}, e...)
		res = append(res, e)
	}

	return res
}

// 724
func pivotIndex(nums []int) int {

	pre := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	for i := 1; i <= len(nums); i++ {
		l, r := pre[i-1], pre[len(nums)]-pre[i]
		if l == r {
			return i - 1
		}
	}
	return -1
}
