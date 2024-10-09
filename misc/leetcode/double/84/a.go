package double_84

import "sort"

// 1
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {

	type node struct {
		val, score int
	}
	h := make(map[int]*node)
	nodeList := make([]*node, 0)
	for _, item := range items1 {
		n, ok := h[item[0]]
		if !ok {
			n = &node{val: item[0], score: item[1]}
			h[item[0]] = n
			nodeList = append(nodeList, n)
		}
	}

	for _, item := range items2 {
		n, ok := h[item[0]]
		if !ok {
			n = &node{val: item[0]}
			h[item[0]] = n
			nodeList = append(nodeList, n)
		}
		n.score += item[1]
	}
	sort.Slice(nodeList, func(i, j int) bool { return nodeList[i].val < nodeList[j].val })
	res := make([][]int, 0, len(nodeList))
	for _, node := range nodeList {
		res = append(res, []int{node.val, node.score})
	}
	return res
}

// 2

func countBadPairs(nums []int) int64 {

	h, n := make(map[int]int), len(nums)
	ans := int64(n * (n - 1) / 2)
	for i, v := range nums {
		t := v - i
		ans -= int64(h[t])
		h[t]++
	}
	return ans
}

// 3

func taskSchedulerII(tasks []int, space int) int64 {

	h := make(map[int]int)
	ans := int64(0)
	for _, v := range tasks {
		last, ok := h[v]
		if ok {
			vql := int64(last + space + 1)
			if vql <= ans {
				ans += 1
			} else {
				ans = vql
			}

			h[v] = int(ans)
			continue
		}
		ans += 1
		h[v] = int(ans)
	}
	return ans
}

// 4
func minimumReplacement(nums []int) int64 {

	ans := int64(0)
	last := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > last {
			k := nums[i] / last
			if nums[i]%last > 0 {
				k = k + 1
			}
			ans += int64(k - 1)
			last = nums[i] / k
		} else {
			last = nums[i]
		}
	}
	return ans
}
