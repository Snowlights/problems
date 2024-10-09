package jiukun

import (
	"sort"
	"strconv"
)

// 1
func numberOfPairs(nums []int) int {
	ans, mod := 0, int(1e9+7)
	h := make(map[int]int)
	for _, v := range nums {
		rv := reserve(v)
		val := rv - v
		ans += h[reserve(v)-v]
		ans %= mod
		h[val]++
	}
	return ans
}

func reserve(num int) int {
	s := []byte(strconv.Itoa(num))
	start, end := 0, len(s)-1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	res, _ := strconv.Atoi(string(s))
	return res
}

// 2
func lakeCount(field []string) int {

	ans := 0
	vis := make([][]bool, len(field))
	for i := range vis {
		vis[i] = make([]bool, len(field[0]))
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 左上
		x, y := i-1, j-1
		if x >= 0 && y >= 0 && !vis[x][y] {
			vis[x][y] = true
			if field[x][y] == 'W' {
				dfs(x, y)
			}
		}
		// 上方
		x, y = i-1, j
		if x >= 0 && y >= 0 && !vis[x][y] {
			vis[x][y] = true
			if field[x][y] == 'W' {
				dfs(x, y)
			}
		}
		// 右上
		x, y = i-1, j+1
		if x >= 0 && y < len(field[0]) && !vis[x][y] {
			vis[x][y] = true
			if field[x][y] == 'W' {
				dfs(x, y)
			}
		}

		// 左
		x, y = i, j-1
		if x >= 0 && y >= 0 && !vis[x][y] {
			vis[x][y] = true
			if field[x][y] == 'W' {
				dfs(x, y)
			}
		}
		// 右
		x, y = i, j+1
		if x >= 0 && y < len(field[0]) && !vis[x][y] {
			vis[x][y] = true
			if field[x][y] == 'W' {
				dfs(x, y)
			}
		}
		// 左下
		x, y = i+1, j-1
		if x < len(field) && y >= 0 && !vis[x][y] {
			vis[x][y] = true
			if field[x][y] == 'W' {
				dfs(x, y)
			}
		}
		// 下
		x, y = i+1, j
		if x < len(field) && y >= 0 && !vis[x][y] {
			vis[x][y] = true
			if field[x][y] == 'W' {
				dfs(x, y)
			}
		}
		// 右下
		x, y = i+1, j+1
		if x < len(field) && y < len(field[0]) && !vis[x][y] {
			vis[x][y] = true
			if field[x][y] == 'W' {
				dfs(x, y)
			}
		}
	}

	for i, v := range field {
		for j, vv := range v {
			if vis[i][j] || vv == '.' {
				continue
			}
			vis[i][j] = true
			dfs(i, j)
			ans++
		}
	}
	return ans
}

// 3
func minOperations(numbers []int) int {
	ans := 0
	base := int64(numbers[0])
	for i := 1; i < len(numbers); i++ {
		g := gcd(base, int64(numbers[i]))
		base *= int64(numbers[i]) / g
	}
	for _, v := range numbers {
		tmp := base / int64(v)
		num := 0
		for tmp != 1 {
			if tmp%2 == 0 {
				tmp /= 2
				num++
				continue
			}
			if tmp%3 == 0 {
				tmp /= 3
				num++
				continue
			}
			break
		}
		if tmp != 1 {
			return -1
		}
		ans += num
	}

	return ans
}

func gcd(a, b int64) int64 {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

// 4

var n, m int
var factor [55]int
var mp = map[int64]float64{}

func hash(nums []int) int64 {
	ans := int64(0)
	for i := 0; i < len(nums); i++ {
		ans = ans*int64(factor[i]+1) + int64(nums[i])
	}
	return ans
}

func dfs(nums []int, cc int) float64 {
	h := int64(hash(nums))
	if v, ok := mp[h]; ok {
		return v
	}
	var ans float64
	var eff int
	if cc < m {
		eff += n - cc
	}
	var j int
	for i := m - cc; i < m; i = j {
		for j = i; j < m; j++ {
			if nums[i] != nums[j] {
				break
			}
		}
		nums[j-1] += 1
		if nums[j-1] <= factor[j-1] {
			eff += j - i
		}
		nums[j-1] -= 1
	}
	if cc < m {
		nums[m-cc-1] = 1
		ans += float64(n-cc) / float64(eff) * dfs(nums, cc+1)
		nums[m-cc-1] = 0
	}
	for i := m - cc; i < m; i = j {
		for j = i; j < m; j++ {
			if nums[i] != nums[j] {
				break
			}
		}
		nums[j-1] += 1
		if nums[j-1] <= factor[j-1] {
			ans += float64(j-i) / float64(eff) * dfs(nums, cc)
		}
		nums[j-1] -= 1
	}
	ans += float64(n) / float64(eff)
	mp[h] = ans
	return ans
}

func chipGame(nums []int, kind int) float64 {
	sort.Ints(nums)
	m, n = len(nums), kind
	mp = make(map[int64]float64)
	for i := 0; i < m; i++ {
		factor[i] = nums[i]
	}

	mp[int64(hash(nums))] = 0
	tmp := make([]int, m)
	return dfs(tmp, 0)
}
