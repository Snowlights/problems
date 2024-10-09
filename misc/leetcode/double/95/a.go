package double_95

import "sort"

// 1
func categorizeBox(length int, width int, height int, mass int) string {
	tiji := length * width * height
	if length >= 1e4 || width >= 1e4 || height >= 1e4 || tiji >= 1e9 {
		if mass >= 100 {
			return "Both"
		}
		return "Bulky"
	} else if mass >= 100 {
		return "Heavy"
	}
	return "Neither"
}

// 2
type DataStream struct {
	val, cnt, k int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		val: value,
		cnt: 0,
		k:   k,
	}
}

func (this *DataStream) Consec(num int) bool {
	if num == this.val {
		this.cnt++
	} else {
		this.cnt = 0
	}
	return this.cnt >= this.k
}

// 3
func xorBeauty(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

// 4
func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	sum := make([]int, n+1) // 前缀和
	for i, x := range stations {
		sum[i+1] = sum[i] + x
	}
	return int64(sort.Search(sum[n]+k, func(minPower int) bool {
		minPower++               // 改为二分最小的不满足要求的值，这样 sort.Search 返回的就是最大的满足要求的值
		diff := make([]int, n+1) // 差分数组
		sumD, need := 0, 0
		for i, d := range diff[:n] {
			sumD += d // 累加差分值
			power := sum[min(i+r+1, n)] - sum[max(i-r, 0)]
			m := minPower - power - sumD
			if m > 0 { // 需要 m 个供电站
				need += m
				if need > k { // 提前退出这样快一些
					return true // 不满足要求
				}
				sumD += m                  // 差分更新
				diff[min(i+r*2+1, n)] -= m // 差分更新
			}
		}
		return false
	}))
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
