package _00_400

import (
	"math/rand"
	"strings"
)

// 376
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}

// 377
func combinationSum4(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

// 380
type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indices[val]; ok {
		return false
	}
	rs.indices[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	id, ok := rs.indices[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1
	rs.nums[id] = rs.nums[last]
	rs.indices[rs.nums[id]] = id
	rs.nums = rs.nums[:last]
	delete(rs.indices, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

// 384
type Solution struct {
	h             map[int]bool
	original, arr []int
}

func Constructor384(nums []int) Solution {
	h := make(map[int]bool)
	for _, v := range nums {
		h[v] = true
	}
	return Solution{h, nums, nums}
}

func (this *Solution) Reset() []int {
	this.arr = this.original
	return this.arr
}

func (this *Solution) Shuffle() []int {
	val := make([]int, 0, len(this.original))
	for v, _ := range this.h {
		val = append(val, v)
	}
	this.arr = val
	return this.arr
}

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
