package double_94

import (
	"sort"
	"strings"
)

// 1
func captureForts(forts []int) int {
	ans := 0
	n := len(forts)
	for i, v := range forts {
		if v != 1 {
			continue
		}
		l, r := i-1, i+1
		for l >= 0 && forts[l] == 0 {
			l--
		}
		if l >= 0 && forts[l] == -1 {
			ans = max(ans, i-l-1)
		}
		for r < n && forts[r] == 0 {
			r++
		}
		if r < n && forts[r] == -1 {
			ans = max(ans, r-i-1)
		}
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 2
func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	pos, neg := make(map[string]bool), make(map[string]bool)
	for _, v := range positive_feedback {
		pos[v] = true
	}
	for _, v := range negative_feedback {
		neg[v] = true
	}
	type stu struct {
		id, score int
	}
	stuList := make([]*stu, len(student_id))
	for i := range stuList {
		s := &stu{
			id:    student_id[i],
			score: 0,
		}
		r := strings.Split(report[i], " ")
		for _, v := range r {
			if pos[v] {
				s.score += 3
			}
			if neg[v] {
				s.score--
			}
		}
		stuList[i] = s
	}

	sort.Slice(stuList, func(i, j int) bool {
		s1, s2 := stuList[i], stuList[j]
		return s1.score > s2.score || s1.score == s2.score && s1.id < s2.id
	})

	ans := make([]int, 0, k)
	for _, v := range stuList[:k] {
		ans = append(ans, v.id)
	}
	return ans
}

// 3
func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {

	lcm := divisor1 / gcd(divisor1, divisor2) * divisor2

	check := func(x int) bool {
		left1 := max(uniqueCnt1-x/divisor2+x/lcm, 0)
		left2 := max(uniqueCnt2-x/divisor1+x/lcm, 0)
		common := x - x/divisor1 - x/divisor2 + x/lcm
		return common >= left1+left2
	}

	l, r := 1, (uniqueCnt1+uniqueCnt2)*2-1
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return int(r)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 4
func countAnagrams(s string) int {

	mod := int(1e9 + 7)
	a, b := 1, 1
	for _, v := range strings.Split(s, " ") {
		cnt := [26]int{}
		for i, c := range v {
			a = a * (i + 1) % mod
			cnt[c-'a']++
			b = b * cnt[c-'a'] % mod
		}
	}

	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	return a * pow(b, mod-2) % mod
}
