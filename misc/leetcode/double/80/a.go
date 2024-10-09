package double_80

import "sort"

// 1
func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	daxie, xiaoxie, shuzi, teshuzifu, lianxu := 0, 0, 0, 0, 0
	for i, v := range password {
		if v >= 'a' && v <= 'z' {
			xiaoxie++
		}
		if v >= 'A' && v <= 'Z' {
			daxie++
		}
		if v >= '0' && v <= '9' {
			shuzi++
		}
		// !@#$%^&*()-+
		if v == '!' || v == '@' || v == '#' ||
			v == '$' || v == '%' || v == '^' ||
			v == '&' || v == '*' || v == '(' || v == ')' || v == '-' || v == '+' {
			teshuzifu++
		}
		if i == 0 {
			continue
		}
		if string(v) == string(password[i-1]) {
			lianxu++
		}
	}
	if daxie > 0 && xiaoxie > 0 && shuzi > 0 && teshuzifu > 0 && lianxu == 0 {
		return true
	}

	return false
}

// 2
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ans := make([]int, len(spells))
	for i, spell := range spells {
		idx := sort.Search(len(potions), func(i int) bool {
			return int64(spell)*int64(potions[i]) >= success
		})
		if idx < len(potions) {
			ans[i] = len(potions) - idx
		} else {
			ans[i] = 0
		}
	}
	return ans
}

// 3

func matchReplacement(s string, sub string, mappings [][]byte) bool {
	m := make(map[byte][]byte)
	for _, v := range mappings {
		m[v[0]] = append(m[v[0]], v[1])
	}
	for i := range s {
		sS, subS := i, 0
		for subS < len(sub) && sS < len(s) {
			sb, subb := s[sS], sub[subS]
			if byte(sb) == subb || canRep(byte(subb), sb, m) {
				subS++
				sS++
				continue
			}
			break
		}
		if subS == len(sub) {
			return true
		}
	}

	return false
}

func canRep(a, b byte, m map[byte][]byte) bool {
	for _, v := range m[a] {
		if v == b {
			return true
		}
	}
	return false
}

// 4
func countSubarrays(nums []int, k int64) (ans int64) {
	sum, left := int64(0), 0
	for right, num := range nums {
		sum += int64(num)
		for sum*int64(right-left+1) >= k {
			sum -= int64(nums[left])
			left++
		}
		ans += int64(right - left + 1)
	}
	return
}
