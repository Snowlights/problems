package _700_1800

import "sort"

// 1757
// select product_id from Products where low_fats = 'Y' and recyclable = 'Y'

// 1760
func minimumSize(nums []int, maxOps int) int {
	// 找最小，往左侧缩小搜索范围
	return sort.Search(1e9, func(i int) bool {
		i++
		tot := 0
		for _, v := range nums {
			// 两种情况:
			// 1. v 能整除 i, 操作次数 = v/i - 1
			// 2. v 不能整除 i, 操作次数 = v/i
			// 合并为: 向下取整
			tot += (v - 1) / i
		}
		return tot <= maxOps
	}) + 1
}
