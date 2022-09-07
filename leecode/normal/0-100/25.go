package __100

import (
	"sort"
	"strings"
)

// 28
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 34
func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	i1 := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if i1 < len(nums) && nums[i1] == target {
		ans[0] = i1
		for i1 < len(nums) && nums[i1] == target {
			i1++
		}
		ans[1] = i1 - 1
	}
	return ans
}

// 35
func searchInsert(nums []int, target int) int {

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

	return sort.Search(len(nums), func(i int) bool { return nums[i] > target })
}

// 42
func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 43
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	ans := make([]int, l1+l2)
	add := func(pos, sum int) {
		for ; sum > 0; pos++ {
			tmp := ans[pos] + sum
			sum = tmp / 10
			ans[pos] = tmp % 10
		}
	}
	p := l1 - 1 + l2 - 1
	for i := l1 - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := l2 - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			add(p-i-j, int(n1*n2))
		}
	}
	k := l1 + l2 - 1
	for ; k >= 0 && ans[k] == 0; k-- {
	}
	var str string
	for i := k; i >= 0; i-- {
		str += string(ans[i] + '0')
	}
	return str
}

// 46
func permute(nums []int) [][]int {

	var bfs func([]int)
	res := make([][]int, 0)
	h := make(map[int]bool)
	bfs = func(ans []int) {
		if len(ans) == len(nums) {
			arr := make([]int, len(ans))
			copy(arr, ans)
			res = append(res, arr)
			return
		}

		for _, v := range nums {
			if h[v] {
				continue
			}
			h[v] = true
			bfs(append(ans, v))
			h[v] = false
		}
	}
	bfs([]int{})
	return res
}

// 48
func rotate(matrix [][]int) {
	arr := round(matrix)
	copy(matrix, arr)
}

func round(mat [][]int) [][]int {
	target := make([][]int, len(mat))

	for i := 0; i < len(mat); i++ {
		target[i] = make([]int, len(mat))
		for j := 0; j < len(mat); j++ {
			target[i][j] = mat[len(mat)-1-j][i]
		}
	}
	return target
}

// 49
func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
