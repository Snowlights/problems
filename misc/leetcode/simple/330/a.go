package simple_330

import "sort"

// 1
func distinctIntegers(n int) int {
	if n == 1 {
		return 1
	}
	return n - 1
}

// 2
func monkeyMove(n int) int {
	mod := int(1e9 + 7)
	pow := func(x, n int) int {
		//x %= mod
		res := int(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	return int(pow(2, int(n))+mod-2) % mod
}

// 3
func putMarbles(wt []int, k int) (ans int64) {
	for i, w := range wt[1:] {
		wt[i] += w
	}
	wt = wt[:len(wt)-1]
	sort.Ints(wt)
	for _, w := range wt[len(wt)-k+1:] {
		ans += int64(w)
	}
	for _, w := range wt[:k-1] {
		ans -= int64(w)
	}
	return
}

// 4 https://leetcode.cn/problems/count-increasing-quadruplets/solution/you-ji-qiao-de-mei-ju-yu-chu-li-pythonja-exja/
func countQuadruplets(nums []int) (ans int64) {
	n := len(nums)
	great := make([][]int, n)
	great[n-1] = make([]int, n+1)
	for k := n - 2; k >= 2; k-- {
		great[k] = append([]int(nil), great[k+1]...)
		for x := nums[k+1] - 1; x > 0; x-- {
			great[k][x]++ // x < nums[k+1]，对于 x，大于它的数的个数 +1
		}
	}

	less := make([]int, n+1)
	for j := 1; j < n-2; j++ {
		for x := nums[j-1] + 1; x <= n; x++ {
			less[x]++ // x > nums[j-1]，对于 x，小于它的数的个数 +1
		}
		for k := j + 1; k < n-1; k++ {
			if nums[j] > nums[k] {
				ans += int64(less[nums[k]] * great[k][nums[j]])
			}
		}
	}
	return
}
