package simple_316

import "sort"

// 1
func haveConflict(event1 []string, event2 []string) bool {
	return (event2[0] <= event1[0] && event1[0] <= event2[1]) ||
		(event2[0] <= event1[1] && event1[1] <= event2[1]) ||
		(event1[0] <= event2[0] && event2[0] <= event1[1]) ||
		(event1[0] <= event2[1] && event2[1] <= event1[1])
}

// 2
func subarrayGCD(nums []int, k int) int {
	ans := 0
	for i, v := range nums {
		if v%k > 0 {
			continue
		}
		for j := i; j < len(nums); j++ {
			v = gcd(nums[j], v)
			if v%k > 0 {
				break
			}
			if v == k {
				ans++
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

// 3
func minCost(nums []int, cost []int) int64 {
	type pair struct {
		num, cost int
	}
	total, note := 0, 0
	pairList := make([]pair, 0)
	for i, v := range nums {
		pairList = append(pairList, pair{v, cost[i]})
		total += cost[i]
	}
	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].num < pairList[j].num
	})

	ans, target := int64(0), 0
	for _, v := range pairList {
		note += v.cost
		if note >= total/2 {
			target = v.num
			break
		}
	}

	for _, v := range pairList {
		ans += int64(abs(v.num-target) * v.cost)
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 4
func makeSimilar(nums []int, target []int) int64 {
	even, odd := []int{}, []int{}
	tEven, tOdd := []int{}, []int{}
	for i, v := range nums {
		if v%2 == 1 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
		if target[i]%2 == 1 {
			tEven = append(tEven, target[i])
		} else {
			tOdd = append(tOdd, target[i])
		}
	}

	sort.Ints(even)
	sort.Ints(tEven)
	sort.Ints(tOdd)
	sort.Ints(odd)
	ans := int64(0)
	for i, v := range even {
		ans += int64(abs(v - tEven[i]))
	}
	for i, v := range odd {
		ans += int64(abs(v - tOdd[i]))
	}

	return ans / 4
}
