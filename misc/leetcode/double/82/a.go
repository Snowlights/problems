package double_82

import "sort"

// 1
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {

	var cal func(node *TreeNode) bool
	cal = func(node *TreeNode) bool {
		if node.Left != nil && node.Right != nil {
			l := cal(node.Left)
			r := cal(node.Right)
			switch node.Val {
			case 2:
				return l || r
			case 3:
				return l && r
			default:
				return false
			}
		}
		return node.Val == 1
	}

	return cal(root)

}

// 2

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	pMap := make(map[int]bool)
	cur, cap := 0, capacity

	for _, v := range passengers {
		pMap[v] = true
	}

	for i, v := range buses {
		cap = capacity
		if i == len(buses)-1 {
			break
		}
		for cap > 0 && cur < len(passengers) && passengers[cur] <= v {
			cap--
			cur++
		}
	}
	cap = capacity
	exitList, exitMap := make([]int, 0), map[int]bool{}
	for cap > 0 && cur < len(passengers) && passengers[cur] <= buses[len(buses)-1] {
		exitList = append(exitList, passengers[cur])
		exitMap[passengers[cur]] = true
		cap--
		cur++
	}
	if len(exitList) == 0 {
		return buses[len(buses)-1]
	}
	max := exitList[len(exitList)-1]
	if cap > 0 && max < buses[len(buses)-1] {
		return buses[len(buses)-1]
	}
	for exitMap[max] || pMap[max] {
		max--
	}
	return max
}

// 3
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {

	h := make(map[int]int)
	vv := -1
	for i, v := range nums1 {
		val := abs(v - nums2[i])
		h[val]++
		if val > vv {
			vv = val
		}
	}

	canHandle := k1 + k2
	for canHandle > 0 && vv > 0 {
		v := h[vv]
		canSub := v
		if canHandle < canSub {
			canSub = canHandle
		}
		h[vv-1] += canSub
		h[vv] -= canSub
		canHandle -= canSub
		vv--
	}

	ans := 0
	for k, v := range h {
		ans += k * k * v
	}
	return int64(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 4 找到每个元素为最小值的左右边界
func validSubarraySize(nums []int, threshold int) int {

	l, r := make([]int, len(nums)), make([]int, len(nums))
	stack := []int{}
	for i, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > v {
			r[stack[len(stack)-1]] = i - 1
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		r[stack[len(stack)-1]] = len(nums) - 1
		stack = stack[:len(stack)-1]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			l[stack[len(stack)-1]] = i + 1
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		l[stack[len(stack)-1]] = 0
		stack = stack[:len(stack)-1]
	}

	for i, v := range nums {
		tmp := r[i] - l[i] + 1
		if tmp*v > threshold {
			return tmp
		}
	}

	return -1
}
