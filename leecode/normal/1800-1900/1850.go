package _800_1900

import "sort"

// 1855
func maxDistance(x, y []int) (ans int) {
	for j, v := range y {
		i := sort.Search(min(len(x), j), func(i int) bool { return x[i] <= v })
		if i < min(len(x), j) && j-i > ans {
			ans = j - i
		}
	}
	return
}

// 1870
func minSpeedOnTime(dist []int, hour float64) int {
	res := sort.Search(1e7+1, func(i int) bool {
		if i == 0 {
			return false
		}
		ans := 0.0
		for j, v := range dist {
			if j == len(dist)-1 {
				ans += float64(v) / float64(i)
				continue
			}
			ans += float64(v / i)
			if v%i > 0 {
				ans++
			}
		}
		return ans <= hour
	})
	if res == 1e7+1 {
		return -1
	}
	return res
}

// 1873
// SELECT
//    employee_id,
// IF(MOD(employee_id,2)!=0 AND LEFT(name,1)!='M',salary,0) bonus
// FROM Employees
// ORDER BY employee_id
