package main

func numberOfSubsequences(nums []int) (ans int64) {
	n := len(nums)
	type pair struct{ x, y int }
	suf := map[pair]int{}
	// 枚举 c
	for i := 4; i < n-2; i++ {
		c := nums[i]
		// 枚举 d
		for _, d := range nums[i+2:] {
			g := gcd(c, d)
			suf[pair{d / g, c / g}]++
		}
	}

	// 枚举 b
	for i := 2; i < n-4; i++ {
		b := nums[i]
		// 枚举 a
		for _, a := range nums[:i-1] {
			g := gcd(a, b)
			ans += int64(suf[pair{a / g, b / g}])
		}
		// 撤销之前统计的 d'/c'
		c := nums[i+2]
		for _, d := range nums[i+4:] {
			g := gcd(c, d)
			suf[pair{d / g, c / g}]--
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func numberOfSubsequences2(nums []int) (ans int64) {
	n := len(nums)
	cnt := map[float32]int{}
	// 枚举 b 和 c
	for i := 4; i < n-2; i++ {
		// 增量式更新，本轮循环只需枚举 b=nums[i-2] 这一个数
		// 至于更前面的 b，已经在前面的循环中添加到 cnt 中了，不能重复添加
		b := float32(nums[i-2])
		// 枚举 a
		for _, a := range nums[:i-3] {
			cnt[float32(a)/b]++
		}

		c := float32(nums[i])
		// 枚举 d
		for _, d := range nums[i+2:] {
			ans += int64(cnt[float32(d)/c])
		}
	}
	return
}
