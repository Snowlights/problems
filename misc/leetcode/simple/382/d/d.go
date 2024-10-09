package main

func minOrAfterOperations(nums []int, k int) (ans int) {
	mask := 0
	for b := 29; b >= 0; b-- {
		mask |= 1 << b
		cnt := 0  // 操作次数
		and := -1 // -1 的二进制全为 1
		for _, x := range nums {
			and &= x & mask
			if and != 0 {
				cnt++ // 合并 x，操作次数加一
			} else {
				and = -1 // 准备合并下一段
			}
		}
		if cnt > k {
			ans |= 1 << b  // 答案的这个比特位必须是 1
			mask ^= 1 << b // 后面不考虑这个比特位
		}
	}
	return
}
