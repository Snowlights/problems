package _00_400

import "strings"

// 392
func isSubsequence(s string, t string) bool {

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		for j < len(t) && s[i] != t[j] {
			j++
		}
		if j == len(t) {
			break
		}
		i++
		j++
	}
	return i == len(s)
}

// 394
func decodeString(s string) string {
	numQ, cQ := make([]int, 0), make([]string, 0)
	num, ans := 0, ""
	for _, v := range s {
		if '0' <= v && v <= '9' {
			num = num*10 + int(v-'0')
		} else if v == '[' {
			numQ = append(numQ, num)
			num = 0
			cQ = append(cQ, ans)
			ans = ""
		} else if v == ']' {
			count := numQ[len(numQ)-1]
			numQ = numQ[:len(numQ)-1]
			str := cQ[len(cQ)-1]
			cQ = cQ[:len(cQ)-1]
			ans = str + strings.Repeat(ans, count)
		} else {
			ans += string(v)
		}
	}

	return ans
}
