package lccup_21_spring

import "math"

// 1
func storeWater(bucket []int, vat []int) int {
	finish := true
	for _, v := range vat {
		if v != 0 {
			finish = false
			break
		}
	}
	if finish {
		return 0
	}

	//需要额外补多少下才能满足能在time次内完成
	ans := int(2 * 1e4)

	for time := 1; time <= int(1e4); time++ {
		finish := true
		add := 0
		for i, _ := range bucket {
			n := int(math.Ceil(float64(vat[i]) / float64(time)))
			if n > bucket[i] {
				add += n - bucket[i]
				finish = false
			}
		}

		if finish {
			return min(ans, time)
		}

		ans = min(time+add, ans)
	}

	return ans

}

// 2
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxValue(root *TreeNode, k int) int {
	var dfs func(root *TreeNode, k int) []int
	dfs = func(root *TreeNode, k int) []int {
		if root == nil {
			return []int{0}
		}
		l, r := dfs(root.Left, k), dfs(root.Right, k)
		res := make([]int, k+1)
		for i := 0; i < len(l); i++ {
			for j := 0; j < len(r); j++ {
				res[0] = max(res[0], l[i]+r[j])
				if i+j+1 <= k {
					res[i+j+1] = max(res[i+j+1], l[i]+r[j]+root.Val)
				}
			}

		}
		return res
	}

	result, ans := dfs(root, k), 0
	for i := 0; i <= len(result); i++ {
		if i <= k {
			ans = max(ans, result[i])
		}
	}
	return ans
}

// 3
// todo 第三题 https://leetcode.cn/contest/season/2021-spring/

// 4

// 5

// 6
