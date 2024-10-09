package simple_306

import (
	"math"
	"strconv"
)

// 1

func largestLocal(grid [][]int) [][]int {

	res := make([][]int, len(grid)-2)
	for i := range res {
		res[i] = make([]int, len(grid)-2)
	}

	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid)-2; j++ {
			m := grid[i][j]
			for k := i; k <= i+2; k++ {
				for l := j; l <= j+2; l++ {
					m = max(m, grid[k][l])
				}
			}
			res[i][j] = m
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 2
func edgeScore(edges []int) int {
	h := make(map[int]int)
	for i, v := range edges {
		h[v] += i
	}
	val, ans := math.MinInt32, math.MaxInt32
	for k, v := range h {
		if v >= val {
			if v == val {
				ans = min(k, ans)
			} else {
				ans = k
			}
			val = v
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 3

// pattern 一定只有两个字符 'I' 和 'D':
// 如果只有 'I', 那么答案一定是 (1,2,3,4,5,6...)
// 那么加上 'D' 会对哪些字符有影响呢?
// 可以发现: 一个'D'或者一堆连续的'D'只会对下一个紧邻'I'的有影响.
// 即, 'III(DI)I' 'III(DDI)I(DDI)I(DI)' 有影响的被括号包围.
// 可以发现一种贪心策略: 把有影响的部分逆续就好了
// 比如: IIIDI, 答案就是 从 (1,2,3,4,5,6) --变为>> (1,2,3,5,4,6)
// solution 1
func smallestNumberS1(pattern string) string {
	n := len(pattern)
	bs := []byte{}
	for i := 0; i <= n; i++ {
		bs = append(bs, byte(i+1+'0'))
	}
	swap := func(s, e int) {
		for i, j := s, e; i < j; i, j = i+1, j-1 {
			bs[i], bs[j] = bs[j], bs[i]
		}
	}
	i := 0
	for i < n {
		j := i
		if pattern[i] == 'D' {
			for j < n && pattern[j] == 'D' {
				j++
			}
		}
		swap(i, j)
		i = j + 1
	}
	return string(bs)
}

// solution 2
func smallestNumberS2(pattern string) string {
	n := len(pattern)
	ans := make([]byte, n+1)
	for i, cur := 0, byte('1'); i < n; {
		if i > 0 && pattern[i] == 'I' {
			i++
		}
		for ; i < n && pattern[i] == 'I'; i++ {
			ans[i] = cur
			cur++
		}
		i0 := i
		for ; i < n && pattern[i] == 'D'; i++ {
		}
		for j := i; j >= i0; j-- {
			ans[j] = cur
			cur++
		}
	}
	return string(ans)
}

// my solution
func smallestNumber(pattern string) string {
	swap := func(ans []byte, start, end int) {
		for start < end {
			ans[start], ans[end] = ans[end], ans[start]
			start++
			end--
		}
	}
	pattern += "I"
	idx, v, ans := 0, '1', make([]byte, len(pattern))
	for i := range ans {
		ans[i] = byte(v)
		v++
	}
	for idx < len(pattern) {
		for idx < len(pattern) && pattern[idx] == 'I' {
			idx++
		}
		start := idx
		for idx < len(pattern) && pattern[idx] == 'D' {
			idx++
		}
		end := idx
		swap(ans, start, end)
	}
	return string(ans)
}

// 4
func countSpecialNumbersS1(n int) (ans int) {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, 1<<10)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(i, mask int, isUpper, isNum bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		if !isLimit && isNum {
			dv := &dp[i][mask]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, mask, false, false)
		}
		d := 1
		if isNum {
			d = 0
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return f(0, 0, true, false)
}

var calMap = map[int]int{
	1: 9,
	2: 9 * 8,
	3: 9 * 8 * 7,
	4: 9 * 8 * 7 * 6,
	5: 9 * 8 * 7 * 6 * 5,
	6: 9 * 8 * 7 * 6 * 5 * 4,
	7: 9 * 8 * 7 * 6 * 5 * 4 * 3,
	8: 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2,
	9: 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1,
}

func countSpecialNumbers(n int) (ans int) {
	if n <= 10 {
		return n
	}

	s := strconv.Itoa(n)
	// 计算位数少于当前的值
	base := 9
	for i := 1; i < len(s); i++ {
		if i == 1 {
			ans += calMap[i]
			continue
		}
		ans += base * calMap[i-1]
	}
	// 算一下开头可以选择多少个
	v, _ := strconv.Atoi(string(s[0]))
	ans += (v - 1) * calMap[len(s)-1]
	// 算确定某一位可以有多少个
	h := make(map[int]bool)
	for i, v := range s {
		tmp := 1
		val, _ := strconv.Atoi(string(v))
		if i == 0 {
			h[val] = true
			continue
		}
		num := val
		for j := 0; j < val; j++ {
			if h[j] {
				num--
			}
		}
		tmp *= num
		base := 10 - len(h) - 1
		for k := i + 1; k < len(s); k++ {
			tmp *= base
			base--
		}
		if len(h) == i {
			ans += tmp
		}
		h[val] = true
	}
	if len(s) == len(h) {
		ans++
	}
	return ans
}
