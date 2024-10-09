package main

func countInterestingSubarrays(nums []int, mod, k int) (ans int64) {
	cnt := map[int]int{0: 1} // 把 s[0]=0 算进去
	s := 0
	for _, x := range nums {
		if x%mod == k {
			s = (s + 1) % mod // 这里取模，下面 cnt[s]++ 就不需要取模了
		}
		ans += int64(cnt[(s-k+mod)%mod]) // 避免减法出现负数
		cnt[s]++
	}
	return
}
