package interview

import (
	"strconv"
)

// 面试题 01.01
func isUnique(s string) bool {
	num := 0
	for _, v := range s {
		if num>>int(byte(v)-'a')&1 == 0 {
			num |= 1 << int(byte(v)-'a')
			continue
		}
		return false
	}
	return true
}

// 面试题 01.02
func canPermutePalindrome(s string) bool {
	h, flag := make(map[int]int), false
	for i := range s {
		h[int(s[i]-'a')]++
	}
	for _, v := range h {
		if v&1 == 1 {
			if flag {
				return false
			}
			flag = true
		}
	}
	return true
}

// 面试题 01.06
func compressString(s string) string {
	if len(s) == 0 {
		return s
	}
	ans := ""
	o, c := s[0], 1
	for i := 1; i < len(s); i++ {
		if s[i] == o {
			c++
			continue
		}
		if s[i] != o {
			ans += string(o) + strconv.Itoa(c)
		}
		o = s[i]
		c = 1
	}
	if c > 0 {
		ans += string(o) + strconv.Itoa(c)
	}
	if len(ans) >= len(s) {
		return s
	}
	return ans
}

// 面试题 01.07
func rotate(matrix [][]int) {
	n := len(matrix)
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr[i][j] = matrix[n-1-j][i]
		}
	}
	for i := range arr {
		for j := range arr[i] {
			matrix[i][j] = arr[i][j]
		}
	}

}
