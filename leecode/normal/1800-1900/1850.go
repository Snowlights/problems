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

// 1873
// SELECT
//    employee_id,
// IF(MOD(employee_id,2)!=0 AND LEFT(name,1)!='M',salary,0) bonus
// FROM Employees
// ORDER BY employee_id
