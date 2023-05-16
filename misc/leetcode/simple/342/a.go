package simple_342

// 1
func findDelayedArrivalTime(arrivalTime, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}

// 2
func s(n, m int) int {
	return (1 + n/m) * (n / m) / 2 * m
}

func sumOfMultiples(n int) int {
	return s(n, 3) + s(n, 5) + s(n, 7) - s(n, 15) - s(n, 21) - s(n, 35) + s(n, 105)
}

// 3
func getSubarrayBeauty(nums []int, k, x int) []int {
	const bias = 50
	cnt := [bias*2 + 1]int{}
	for _, num := range nums[:k-1] { // 先往窗口内添加 k-1 个数
		cnt[num+bias]++
	}
	ans := make([]int, len(nums)-k+1)
	for i, num := range nums[k-1:] {
		cnt[num+bias]++ // 进入窗口（保证窗口有恰好 k 个数）
		left := x
		for j, c := range cnt[:bias] { // 暴力枚举负数范围 [-50,-1]
			left -= c
			if left <= 0 { // 找到美丽值
				ans[i] = j - bias
				break
			}
		}
		cnt[nums[i]+bias]-- // 离开窗口
	}
	return ans
}

// 4
func minOperations(nums []int) int {
	n, gcdAll, cnt1 := len(nums), 0, 0
	for _, x := range nums {
		gcdAll = gcd(gcdAll, x)
		if x == 1 {
			cnt1++
		}
	}
	if gcdAll > 1 {
		return -1
	}
	if cnt1 > 0 {
		return n - cnt1
	}

	minSize := n
	for i := range nums {
		g := 0
		for j, x := range nums[i:] {
			g = gcd(g, x)
			if g == 1 {
				// 这里本来是 j+1，把 +1 提出来合并到 return 中
				minSize = min(minSize, j)
				break
			}
		}
	}
	return minSize + n - 1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
