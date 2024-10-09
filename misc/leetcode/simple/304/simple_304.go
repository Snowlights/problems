package simple_304

import "sort"

// 1
func minimumOperations(nums []int) int {
	ans := 0
	next := make([]int, 0)
	for _, v := range nums {
		if v == 0 {
			continue
		}
		next = append(next, v)
	}
	sort.Ints(next)

	for len(next) > 0 {
		n := make([]int, 0)
		v := next[0]
		for i := 1; i < len(next); i++ {
			if next[i]-v == 0 {
				continue
			}
			n = append(n, next[i]-v)
		}
		next = n
		ans++
	}

	return ans
}

// 2
func maximumGroups(grades []int) int {

	var i int
	for i = 0; i <= len(grades); i++ {
		if i*(i+1) > 2*len(grades) {
			break
		}
	}
	return i - 1
}

// 3
func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)

	cal := func(node int) []int {
		r := make([]int, n)
		for i := 0; i < n; i++ {
			r[i] = n
		}
		for d := 0; node >= 0 && r[node] == n; node = edges[node] {
			r[node] = d
			d++
		}
		return r
	}

	aR, bR := cal(node1), cal(node2)
	ans, minD := -1, n

	for i, v := range aR {
		if bR[i] > v {
			v = bR[i]
		}
		if v < minD {
			minD, ans = v, i
		}
	}
	return ans
}

// 4
// 时间戳解决方式
// 具体来说，初始时间戳 {clock}=1，首次访问一个点 x 时，
// 记录访问这个点的时间 time[x]=clock，然后将 clock 加一。
//
//如果首次访问一个点，则记录当前时间 {startTime}={clock}
// 并尝试从这个点出发，看能否找到环。如果找到了一个之前访问过的点 x，
// 且之前访问 x 的时间不早于 startTime，则说明我们找到了一个新的环，
// 此时的环长就是前后两次访问 x 的时间差，即 {clock}-{time} [x] = clock−time[x]。
//
//取所有环长的最大值作为答案。若没有找到环，则返回 -1−1。

func longestCycle(edges []int) int {

	time := make([]int, len(edges))
	clock, ans := 1, -1
	for x, t := range time {
		if t > 0 {
			continue
		}

		for startTime := clock; x >= 0; x = edges[x] {
			if time[x] > 0 { // 重复访问
				if time[x] >= startTime { // 找到了一个新的环
					ans = max(ans, clock-time[x])
				}
				break
			}
			time[x] = clock
			clock++
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

// 4 拓扑排序
// 先查询入度不为一的节点，然后向下查找

func longestCycleV2(edges []int) int {

	ans := -1
	inDegree := make([]int, len(edges))
	for _, edge := range edges {
		if edge == -1 {
			continue
		}
		inDegree[edge]++
	}
	queue := make([]int, 0)
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		to := edges[v]
		if to == -1 {
			continue
		}

		if inDegree[to]--; inDegree[to] == 0 {
			queue = append(queue, to)
		}
	}

	for k, v := range inDegree {
		if v == 0 {
			continue
		}

		length, tmp, val := 1, k, edges[k]
		for ; val >= 0; val = edges[val] {
			if val == tmp {
				break
			}
			inDegree[val]--
			length++
		}
		ans = max(ans, length)
	}

	return ans
}
