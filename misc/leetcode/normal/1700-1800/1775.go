package _700_1800

import "math/bits"

// 1799
// 我们可以先预处理得到数组 nums 中任意两个数的最大公约数，存储在二维数组 gg 中，
//其中 g[i][j] 表示 nums[i] 和 nums[j] 的最大公约数。
//
//然后定义 f[k]f[k] 表示当前操作后的状态为 kk 时，可以获得的最大分数和。假设 mm 为数组 nums 中的元素个数，那么状态一共有 2^m2
//m 种，即 kk 的取值范围为 [0, 2^m - 1]
//
//从小到大枚举所有状态，对于每个状态 kk，先判断此状态中二进制位的个数 cnt 是否为偶数，是则进行如下操作：
//
//枚举 kk 中二进制位为 1 的位置，假设为 ii 和 jj，则 ii 和 jj 两个位置的元素可以进行一次操作，
//此时可以获得的分数为
// cnt / 2 × g[i][j]，更新 f[k] 的最大值。
//
//最终答案即为 f[2^m - 1]

func maxScore(nums []int) int {
	m := len(nums)
	g := [14][14]int{}
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			g[i][j] = gcd(nums[i], nums[j])
		}
	}
	f := make([]int, 1<<m)
	for k := 0; k < 1<<m; k++ {
		cnt := bits.OnesCount(uint(k))
		if cnt%2 == 0 {
			for i := 0; i < m; i++ {
				if k>>i&1 == 1 {
					for j := i + 1; j < m; j++ {
						if k>>j&1 == 1 {
							f[k] = max(f[k], f[k^(1<<i)^(1<<j)]+cnt/2*g[i][j])
						}
					}
				}
			}
		}
	}
	return f[1<<m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
