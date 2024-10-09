package algorithm

import "fmt"

// 匈牙利算法-最大匹配问题

func hungarian() {
	var n, v int
	even, odd := make([]int, 0), make([]int, 0)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		if v&1 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	isPrime := func(v int) bool {
		for i := 2; i*i <= v; i++ {
			if v%i == 0 {
				return false
			}
		}
		return true
	}

	vis := make(map[int]bool)
	matched := make(map[int]int)
	var match func(odd int) bool
	match = func(odd int) bool {
		for _, e := range even {
			if !vis[e] && isPrime(odd+e) {
				vis[e] = true
				if matched[e] == 0 || match(matched[e]) {
					matched[e] = odd
					return true
				}
			}
		}
		return false
	}

	ans := 0
	for _, v := range odd {
		vis = make(map[int]bool)
		if match(v) {
			ans++
		}
	}

	fmt.Println(ans)
}
