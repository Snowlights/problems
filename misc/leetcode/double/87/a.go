package double_87

import "sort"

// 1
// var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
//
//func calcDays(s string) (day int) {
//	for _, d := range days[:s[0]&15*10+s[1]&15-1] {
//		day += d
//	}
//	return day + int(s[3]&15*10+s[4]&15)
//}
//
//func countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob string) int {
//	ans := calcDays(min(leaveAlice, leaveBob)) - calcDays(max(arriveAlice, arriveBob)) + 1
//	if ans < 0 {
//		ans = 0
//	}
//	return ans
//}
//
//func min(a, b string) string { if b < a { return b }; return a }
//func max(a, b string) string { if b > a { return b }; return a }

// 2
func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	j, m := 0, len(trainers)
	for i, p := range players {
		for j < m && trainers[j] < p {
			j++
		}
		if j == m { // 无法找到匹配的训练师
			return i
		}
		j++ // 匹配一位训练师
	}
	return len(players) // 所有运动员都能匹配
}

// 3
func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	type pair struct{ or, i int }
	ors := []pair{} // 按位或的值 + 对应子数组的右端点的最小值
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		ors = append(ors, pair{0, i})
		ors[0].or |= num
		k := 0
		for _, p := range ors[1:] {
			p.or |= num
			if ors[k].or == p.or {
				ors[k].i = p.i // 合并相同值，下标取最小的
			} else {
				k++
				ors[k] = p
			}
		}
		ors = ors[:k+1]
		// 本题只用到了 ors[0]，如果题目改成任意给定数字，可以在 ors 中查找
		ans[i] = ors[0].i - i + 1
	}
	return ans
}

// 4
func minimumMoney(transactions [][]int) int64 {
	totalLose, mx := 0, 0
	for _, t := range transactions {
		totalLose += max(t[0]-t[1], 0)
		mx = max(mx, min(t[0], t[1]))
	}
	return int64(totalLose + mx)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
