package _800_1900

import (
	"container/heap"
	"math"
	"sort"
)

// 1880
func isSumEqual(firstWord, secondWord, targetWord string) bool {
	return str2int(firstWord)+str2int(secondWord) == str2int(targetWord)
}

func str2int(s string) (x int) {
	for _, b := range s {
		x = x*10 + int(b-'a')
	}
	return
}

// 1881
func maxValue(n string, x int) string {
	b := []byte(n)
	if n[0] == '-' {
		for i, v := range n {
			if i == 0 {
				continue
			}
			if x < int(v-'0') {
				return string(append(b[:i], append([]byte{byte(x + '0')}, b[i:]...)...))
			}
		}
		return n + string(rune(x+'0'))
	} else {
		for i, v := range n {
			if x > int(v-'0') {
				return string(append(b[:i], append([]byte{byte(x + '0')}, b[i:]...)...))
			}
		}
		return n + string(rune(x+'0'))
	}
}

// 1882
type Server struct {
	weight int
	index  int
}

// 服务器队列
type server []Server

func (s server) Len() int      { return len(s) }
func (s server) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s server) Less(i, j int) bool {
	return s[i].weight < s[j].weight || s[i].weight == s[j].weight && s[i].index < s[j].index
}
func (s *server) Push(val interface{}) { *s = append(*s, val.(Server)) }
func (s *server) Pop() interface{}     { ret := (*s)[len(*s)-1]; *s = (*s)[:len(*s)-1]; return ret }

// 执行队列
type Exec struct {
	endTime int
	server  Server
}

type ExecQ []Exec

func (e ExecQ) Len() int      { return len(e) }
func (e ExecQ) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e ExecQ) Less(i, j int) bool {
	return e[i].endTime < e[j].endTime || e[i].endTime == e[j].endTime && e[i].server.weight < e[j].server.weight || e[i].endTime == e[j].endTime && e[i].server.weight == e[j].server.weight && e[i].server.index < e[j].server.index
}
func (e *ExecQ) Push(val interface{}) { *e = append(*e, val.(Exec)) }
func (e *ExecQ) Pop() interface{}     { ret := (*e)[len(*e)-1]; *e = (*e)[:len(*e)-1]; return ret }

func assignTasks(servers []int, tasks []int) []int {
	m := len(tasks)
	ans := make([]int, m)
	s := server(nil)
	execQ := ExecQ(nil)
	for i, v := range servers {
		elem := Server{weight: v, index: i}
		heap.Push(&s, elem)
	}

	for i, v := range tasks {
		for len(execQ) > 0 && execQ[0].endTime <= i {
			toe := heap.Pop(&execQ).(Exec)
			heap.Push(&s, toe.server)
		}
		if len(s) != 0 {
			tos := heap.Pop(&s).(Server)
			ans[i] = tos.index
			heap.Push(&execQ, Exec{endTime: i + v, server: tos})
		} else {
			toe := heap.Pop(&execQ).(Exec)
			ans[i] = toe.server.index
			heap.Push(&execQ, Exec{endTime: toe.endTime + v, server: toe.server})
		}
	}
	return ans
}

// 1884
func minSkips(dist []int, speed, hoursBefore int) (ans int) {
	n := len(dist)
	ceilDist := func(d int) int { return (d + speed - 1) / speed * speed }
	f := make([]int, n)
	for i := range f {
		f[i] = 2e9
	}
	f[0] = 0
	for _, d := range dist {
		for j := n - 1; j > 0; j-- {
			f[j] = min(ceilDist(f[j]+d), f[j-1]+d)
		}
		f[0] += ceilDist(d)
	}
	for i, d := range f { // 从小到大遍历跳过次数
		if (d+speed-1)/speed <= hoursBefore {
			return i // 返回第一个满足时间要求的
		}
	}
	return -1
}

// 1885
func twoEggDrop(n int) int {
	return sort.Search(n, func(i int) bool {
		return (1+i)*i >= n*2
	})
}

// 1886
func findRotation(mat [][]int, target [][]int) bool {

	if equals(mat, target) {
		return true
	}

	for i := 0; i < 3; i++ {
		mat = round(mat)
		if equals(target, mat) {
			return true
		}
	}
	return false
}

func equals(mat, target [][]int) bool {
	equal := true
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if mat[i][j] != target[i][j] {
				equal = false
				break
			}
		}
		if !equal {
			break
		}
	}
	return equal
}

func round(mat [][]int) [][]int {
	target := make([][]int, len(mat))

	for i := 0; i < len(mat); i++ {
		target[i] = make([]int, len(mat))
		for j := 0; j < len(mat); j++ {
			target[i][j] = mat[len(mat)-1-j][i]
		}
	}
	return target
}

// 1887
func reductionOperations(nums []int) int {

	type node struct {
		val, num int
	}
	nodeList, nodeMap := make([]*node, 0), make(map[int]*node)
	for _, v := range nums {
		n, ok := nodeMap[v]
		if !ok {
			n = &node{val: v}
			nodeList = append(nodeList, n)
			nodeMap[v] = n
		}
		n.num++
	}

	sort.Slice(nodeList, func(i, j int) bool { return nodeMap[i].val > nodeMap[j].val })
	ans := 0
	for i := 0; i <= len(nodeList)-1; i++ {
		ans += nodeList[i].num
		nodeList[i+1].num += nodeList[i].num
	}
	return ans
}

// 1888
func minFlips(s string) int {
	n := len(s)
	ans := n
	// 枚举开头是 0 还是 1
	for head := byte('0'); head <= '1'; head++ {
		// 左边每个位置的不同字母个数
		leftDiff := make([]int, n)
		diff := 0
		for i := range s {
			if s[i] != head^byte(i&1) {
				diff++
			}
			leftDiff[i] = diff
		}

		// 右边每个位置的不同字母个数
		tail := head ^ 1
		diff = 0
		for i := n - 1; i >= 0; i-- {
			// 左边+右边即为整个字符串的不同字母个数，取最小值
			ans = min(ans, leftDiff[i]+diff)
			if s[i] != tail^byte((n-1-i)&1) {
				diff++
			}
		}
	}
	return ans
}

// 1889
// 由于包裹的尺寸之和是固定的，因此目标转换为最小化箱子的尺寸之和。
// 将包裹按尺寸排序后，按照箱子尺寸划分包裹，使得每个包裹都装入能装该包裹的最小的箱子。
// 这种贪心策略能使箱子的尺寸总和尽可能地小
func minWastedSpace(a []int, boxes [][]int) int {
	sort.Ints(a) // 将包裹按尺寸排序，方便后面按照箱子尺寸划分包裹
	ans := math.MaxInt64
	for _, box := range boxes {
		sort.Ints(box)
		if box[len(box)-1] < a[len(a)-1] { // 最大的箱子不够装最大的包裹
			continue
		}
		res, l := 0, 0
		for _, v := range box {
			// 划分包裹：当前箱子 v 可以装入下标在 [l, r) 区间内的包裹
			r := sort.SearchInts(a, v+1)
			res += (r - l) * v // 统计箱子尺寸之和
			l = r
		}
		ans = min(ans, res)
	}
	if ans < math.MaxInt64 {
		for _, v := range a {
			ans -= v // 减去每个包裹尺寸
		}
		return ans % (1e9 + 7)
	}
	return -1
}

// 1890

func F1890() {
	// SELECT user_id, max(time_stamp) last_stamp  #求最大的日期用max，但是如何限定是2020年呢？
	// FROM Logins
	// WHERE year(time_stamp) = '2020'  #看这！！！！！！！用year函数增加条件为2020年
	// GROUP BY user_id;               #这个好理解就是分个组
	//
}

// 1893
func isCovered(ranges [][]int, left, right int) bool {
	diff := [52]int{} // 差分数组
	for _, r := range ranges {
		diff[r[0]]++
		diff[r[1]+1]--
	}
	cnt := 0 // 前缀和
	for i := 1; i <= 50; i++ {
		cnt += diff[i]
		if cnt <= 0 && left <= i && i <= right {
			return false
		}
	}
	return true
}

// 1894
func chalkReplacer(chalk []int, k int) int {
	total := 0
	for _, v := range chalk {
		total += v
	}
	k %= total

	for {
		for i, v := range chalk {
			if k < v {
				return i
			}
			k -= v
		}
	}
	return -1
}

// 1895
func largestMagicSquare(a [][]int) int {
	n, m := len(a), len(a[0])
	rs := make([][]int, n)
	cs := make([][]int, n+1)
	ds := make([][]int, n+1)
	as := make([][]int, n+1)
	for i := range cs {
		cs[i] = make([]int, m)
		ds[i] = make([]int, m+1)
		as[i] = make([]int, m+1)
	}
	for i, row := range a {
		rs[i] = make([]int, m+1)
		for j, v := range row {
			rs[i][j+1] = rs[i][j] + v
			cs[i+1][j] = cs[i][j] + v
			ds[i+1][j+1] = ds[i][j] + v
			as[i+1][j] = as[i][j+1] + v
		}
	}

	for k := min(n, m); ; k-- {
		for r := k; r <= n; r++ {
		outer:
			for c := k; c <= m; c++ {
				s := ds[r][c] - ds[r-k][c-k]
				if as[r][c-k]-as[r-k][c] != s {
					continue
				}
				for _, row := range rs[r-k : r] {
					if row[c]-row[c-k] != s {
						continue outer
					}
				}
				for j := c - k; j < c; j++ {
					if cs[r][j]-cs[r-k][j] != s {
						continue outer
					}
				}
				return k
			}
		}
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 1897
func makeEqual(a []string) bool {
	c := [26]int{}
	for _, s := range a {
		for _, b := range s {
			c[b-'a']++
		}
	}
	for _, v := range c {
		if v%len(a) > 0 {
			return false
		}
	}
	return true
}

// 1898
func maximumRemovals(s, p string, id []int) int {
	return sort.Search(len(id), func(k int) bool {
		del := make([]bool, len(s))
		for _, i := range id[:k+1] {
			del[i] = true
		}
		j := 0
		for i := range s {
			if !del[i] && s[i] == p[j] {
				if j++; j == len(p) {
					return false
				}
			}
		}
		return true
	})
}

// 1899
func mergeTriplets(a [][]int, t []int) bool {
	found := [3]bool{}
	for _, p := range a {
		if p[0] <= t[0] && p[1] <= t[1] && p[2] <= t[2] {
			found[0] = found[0] || p[0] == t[0]
			found[1] = found[1] || p[1] == t[1]
			found[2] = found[2] || p[2] == t[2]
		}
	}
	return found[0] && found[1] && found[2]
}
