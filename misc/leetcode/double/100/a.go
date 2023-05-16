package double_100

import (
	"math"
	"sort"
)

// 1
func distMoney(money, children int) int {
	money -= children // 每人至少 1 美元
	if money < 0 {
		return -1
	}
	ans := min(money/7, children) // 初步分配，让尽量多的人分到 8 美元
	money -= ans * 7
	children -= ans
	if children == 0 && money > 0 || // 必须找一个前面分了 8 美元的人，分完剩余的钱
		children == 1 && money == 3 { // 不能有人恰好分到 4 美元
		ans--
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 2
func maximizeGreatness(nums []int) int {
	maxCnt := 0
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
		maxCnt = max(maxCnt, cnt[v])
	}
	return len(nums) - maxCnt
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 3
func findScore(nums []int) (ans int64) {
	for i, n := 0, len(nums); i < n; i += 2 { // i 选了 i+1 不能选
		i0 := i
		for i+1 < n && nums[i] > nums[i+1] { // 找到下坡的坡底
			i++
		}
		for j := i; j >= i0; j -= 2 { // 从坡底 i 到坡顶 i0，每隔一个累加
			ans += int64(nums[j])
		}
	}
	return
}

// 4
func repairCars(ranks []int, cars int) int64 {
	minR := ranks[0]
	cnt := [101]int{}
	for _, r := range ranks {
		if r < minR {
			minR = r
		}
		cnt[r]++
	}
	return int64(sort.Search(minR*cars*cars, func(t int) bool {
		s := 0
		for r := minR; r <= 100; r++ {
			s += int(math.Sqrt(float64(t/r))) * cnt[r]
		}
		return s >= cars
	}))
}
