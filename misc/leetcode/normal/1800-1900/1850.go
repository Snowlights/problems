package _800_1900

import (
	"fmt"
	"sort"
)

// 1850
func getMinSwaps(num string, k int) int {

	original, ans := []byte(num), []byte(num)
	for i := 0; i < k; i++ {
		ans = next(ans)
	}
	fmt.Println(string(original), string(ans))
	n, res := len(original), 0
	for i := range original {
		if original[i] == ans[i] {
			continue
		}

		j := i + 1
		for j < n && original[j] != ans[i] {
			j++
		}

		for ; j > i+1; j-- {
			res++
			original[j], original[j-1] = original[j-1], original[j]
		}
	}

	return res
}

func next(original []byte) []byte {
	n := len(original)

	l := n - 2
	for l >= 0 && original[l] >= original[l+1] {
		l--
	}

	if l >= 0 {
		r := n - 1
		for r >= 0 && original[l] >= original[r] {
			r--
		}
		original[l], original[r] = original[r], original[l]
	}

	l++
	n--
	for l < n {
		original[l], original[n] = original[n], original[l]
		l++
		n--
	}

	return original
}

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
