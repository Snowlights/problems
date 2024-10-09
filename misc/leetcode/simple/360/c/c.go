package main

import "math/bits"

func minOperations(nums []int, target int) (ans int) {
	s := 0
	cnt := [31]int{}
	for _, v := range nums {
		s += v
		cnt[bits.TrailingZeros(uint(v))]++
	}
	if s < target {
		return -1
	}
	s = 0
	for i := 0; i < 31; i++ {
		s += cnt[i] << i
		mask := 1<<(i+1) - 1
		if s < target&mask {
			for j := i + 1; ; j++ {
				if cnt[j] > 0 {
					ans += j - i
					cnt[j]--
					s += 1 << j // 从 i 到 j 都至少有一个 1，可以提前加到 s 中
					break
				}
			}
		}
	}
	return
}
