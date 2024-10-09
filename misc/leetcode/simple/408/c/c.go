package main

func numberOfSubstrings(s string) (ans int) {
	a := []int{}
	for i, b := range s {
		if b == '0' {
			a = append(a, i)
		}
	}

	n := len(s)
	tot1 := n - len(a)
	a = append(a, n) // 哨兵

	for left, b := range s {
		if b == '1' {
			ans += a[0] - left // 不含 0 的子串个数
		}
		for k, j := range a[:len(a)-1] {
			cnt0 := k + 1
			if cnt0*cnt0 > tot1 {
				break
			}
			cnt1 := j - left - k
			ans += max(a[k+1]-j-max(cnt0*cnt0-cnt1, 0), 0)
		}
		if b == '0' {
			a = a[1:] // 这个 0 后面不会再枚举到了
		}
	}
	return
}

func numberOfSubstrings2(s string) (ans int) {
	tot1 := 0
	a := []int{-1} // 哨兵
	for right, b := range s {
		if b == '0' {
			a = append(a, right)
		} else {
			ans += right - a[len(a)-1]
			tot1++
		}
		for k := len(a) - 1; k > 0 && (len(a)-k)*(len(a)-k) <= tot1; k-- {
			cnt0 := len(a) - k
			cnt1 := right - a[k] + 1 - cnt0
			ans += max(a[k]-a[k-1]-max(cnt0*cnt0-cnt1, 0), 0)
		}
	}
	return
}
