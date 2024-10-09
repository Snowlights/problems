package didi

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
)

func LCDiDi2019001(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var np, nq, nr int
	fmt.Fscan(r, &np, &nq, &nr)
	type pair struct {
		pre, p, q, r int
	}
	h := make(map[pair]int)

	var dfs func(int, int, int, int) int
	dfs = func(pre, p, q, r int) int {
		pa := pair{pre, p, q, r}
		if val, ok := h[pa]; ok {
			return val
		}

		if p+q+r == 0 {
			return 1
		}

		res := 0
		if pre != 0 && p > 0 {
			res += dfs(0, p-1, q, r)
		}
		if pre != 1 && q > 0 {
			res += dfs(1, p, q-1, r)
		}
		if pre != 2 && r > 0 {
			res += dfs(2, p, q, r-1)
		}
		h[pa] = res
		return res
	}

	res := 0
	if np > 0 {
		res += dfs(0, np-1, nq, nr)
	}
	if nq > 0 {
		res += dfs(1, np, nq-1, nr)
	}
	if nr > 0 {
		res += dfs(2, np, nq, nr-1)
	}
	fmt.Fprintln(w, res)
}

func LCDiDi2019002(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	s1, s2 := "qwertasdfgzxcv", "yuiophjklbnm"
	h1, h2 := make(map[byte]bool), make(map[byte]bool)
	for i := range s1 {
		h1[s1[i]] = true
	}
	for i := range s2 {
		h2[s2[i]] = true
	}

	minDis := func(a, b string) int {
		m, n := len(a), len(b)
		if m*n == 0 {
			return m + n
		}
		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n+1)
		}
		for i := 0; i <= m; i++ {
			dp[i][0] = 3 * i
		}
		for i := 0; i <= n; i++ {
			dp[0][i] = 3 * i
		}
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if a[i-1] == b[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 3
					// 替换
					if (h1[a[i-1]] && h1[b[j-1]]) || (h2[a[i-1]] && h2[b[j-1]]) {
						dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
					} else if (h1[a[i-1]] && h2[b[j-1]]) || (h2[a[i-1]] && h1[b[j-1]]) {
						dp[i][j] = min(dp[i][j], dp[i-1][j-1]+2)
					}
				}
			}
		}
		return dp[m][n]
	}

	s, _, _ := r.ReadLine()
	parts := strings.Split(string(s), " ")
	target := parts[0]
	words := parts[1:]
	target = strings.ToLower(target)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	sort.Slice(words, func(i, j int) bool {
		return minDis(target, words[i]) < minDis(target, words[j])
	})

	for i := 0; i < 3 && i < len(words); i++ {
		fmt.Fprint(w, words[i], " ")
	}

}

func LCDiDi2019003(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	arr := make([]int, n)
	mmin, sum := math.MaxInt32, 0
	for i := range arr {
		fmt.Fscan(r, &arr[i])
		mmin = min(mmin, arr[i])
		sum += arr[i]
	}

	check := func(i int) bool {
		ans, sum := 0, 0
		for _, v := range arr {
			sum += v
			if sum >= i {
				sum = 0
				ans++
			}
		}
		return n-ans <= m
	}

	left, right := mmin, sum
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(w, right)

}

func LCDiDi2019004(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	q, i := []int{}, 1
	var val int
	for j := 0; j < 4; j++ {
		fmt.Fscan(r, &val)
		for {
			if len(q) > 0 && q[len(q)-1] == val {
				q = q[:len(q)-1]
				break
			}
			if i >= 5 {
				break
			}
			q = append(q, i)
			i++
		}
	}

	if len(q) > 0 {
		fmt.Fprintln(w, "No")
	} else {
		fmt.Fprintln(w, "Yes")
	}

}

func LCDiDi2019005(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	// 输入：
	//- 第一行一个整数 n，表示共 n 个城市
	//- 第二行一个整数 m，表示共 m 个航空公司
	//- 第三行和第四行分别表示后面表示航班的二维数组的行数和列数，其中行数表示共有多少航班，列数固定为 4
	//- 从第五行开始，每行四个整数“f t a p”，用空格分隔，f 和 t 分别表示出发城市和到达城市的编号，
	// 取值范围 0 到 n-1，其中 0 表示 A 市，1 表示 B 市;
	// a 表示航空公司编号，取值范围 0 到 m-1;p 表示价格，取值范围大于 0
	//输出：
	//- 一个整数，表示最便宜的机票购买方案的总价格
	//

	var n, m, row, col int
	fmt.Fscan(r, &n, &m, &row, &col)
	var f, t, a, p int

	company := make(map[int]map[int][][]int)
	cost := make([][]float64, n)
	for i := range cost {
		cost[i] = make([]float64, n)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < row; i++ {
		fmt.Fscan(r, &f, &t, &a, &p)
		if company[a] == nil {
			company[a] = make(map[int][][]int)
		}
		company[a][f] = append(company[a][f], []int{t, p})
		cost[f][t] = minFloat(cost[f][t], float64(p))
	}

	for _, info := range company {
		for f := 0; f < n; f++ {
			val, ok := info[f]
			if !ok {
				continue
			}
			for _, to := range val {
				for _, to1 := range info[to[0]] {
					cost[f][to1[0]] = minFloat(cost[f][to1[0]], float64(to[1]+to1[1])*0.7)
				}
			}
		}
	}

	dis := make([]float64, n)
	for i := range dis {
		dis[i] = math.MaxInt32
	}
	dis[0] = 0
	h := &hp{}
	heap.Init(h)
	heap.Push(h, pair{
		cost: 0,
		to:   0,
	})
	hash := make(map[int]bool)
	for i := 0; i < n; i++ {
		hash[i] = true
	}

	for len(hash) > 0 {
		p := heap.Pop(h).(pair)
		if p.to == 1 {
			break
		}
		if !hash[p.to] {
			continue
		}
		for j := range cost[p.to] {
			if len(hash) > 0 {
				if p.cost+cost[p.to][j] < dis[j] {
					dis[j] = float64(p.cost) + cost[p.to][j]
					heap.Push(h, pair{
						cost: dis[j],
						to:   j,
					})
				}
			}
		}
	}
	fmt.Fprintln(w, dis[1])
}

type pair struct {
	cost float64
	to   int
}
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].cost < h[j].cost || (h[i].cost == h[j].cost && h[i].to < h[j].to)
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func LCDiDi201906(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n, d int
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n, &d)
		// 建立dp，用于存储下一天是活动2，且第 j+1 天是活动1，
		// 之前有 i+1 天活动1的值保存为 dp[i][j]
		dp := make([][]int, n/2)
		for i := range dp {
			dp[i] = make([]int, n)
		}
		// 初始化假期开始k天活动1
		for i := 0; i < d; i++ {
			dp[i][i] = 1
		}

		for i := 0; i < n/2; i++ {
			for j := 0; j < n; j++ {
				// 与下一次活动1间间隔的活动2天数 k+1
				for k := 0; k < d; k++ {
					// 下一次活动连续的活动1天数 r+1
					for r := 0; r < d; r++ {
						if i+1+r < n/2 && j+k+2+r < n {
							dp[i+1+r][j+k+2+r] += dp[i][j]
						}
					}
				}
			}
		}

		ans := 0
		// 判断倒数第二天前活动1天数达到n/2的数量
		for i, v := range dp[n/2-1] {
			if i == len(dp[n/2-1])-1 {
				continue
			}
			ans += v
		}
		// 活动1和2交换依然成立，结果乘以2
		fmt.Fprintln(w, ans*2)
	}

}

func LCDiDi201909(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n int
	var op string
	var x int

	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n)
		val := make([]int, n+1)
		for i := range val {
			val[i] = 1
		}
		LineList, id := make([]*Line, 0), 0
		u := newUnionFind(n+1, val)

		for i := 0; i < n; i++ {
			fmt.Fscan(r, &op)
			switch op {
			case "T":
				l := &Line{id: id}
				id++
				fmt.Fscan(r, &l.x1, &l.y1, &l.x2, &l.y2)
				for _, line := range LineList {
					if check(line, l) {
						u.merge(l.id, line.id)
					}
				}
				LineList = append(LineList, l)
			case "Q":
				fmt.Fscan(r, &x)
				fmt.Fprintln(w, u.size(x-1))
			}
		}

		fmt.Fprintln(w)
	}

}

func LCDiDi2019010(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	// 找到在密文各个数字串中都没有出现过的最短子串，若有多个，选字典序最小的
	// 输入：
	//    2
	//    9527
	//    0012345678
	// 输出：02
	var n int
	fmt.Fscan(r, &n)
	arr := make([]string, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}

	cur, l := 0, 1
	for {
		c := fmt.Sprintf("%d", cur)
		x := strings.Repeat("0", l-len(c)) + c
		flag := false
		for _, v := range arr {
			if strings.Contains(v, x) {
				flag = true
				break
			}
		}
		if !flag {
			fmt.Fprintln(w, x)
			break
		}
		if c == strings.Repeat("9", l) {
			l++
			cur = 0
		} else {
			cur++
		}
	}

}

func LCDiDi2019011(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var x1, y1, x2, y2 int
	fmt.Fscan(r, &x1, &y1, &x2, &y2)

	if x1 > y1 || x2 > y2 {
		fmt.Fprintln(w, 0)
		return
	}
	mod := int(1e9 + 7)
	var comb func(n, m int) int
	var fastPower func(x int, n int) int

	fastPower = func(x int, n int) int {
		if n == 0 {
			return 1
		}
		if n%2 == 0 {
			return fastPower((x*x)%mod, n>>1)
		}
		return (x * fastPower((x*x)%mod, n>>1)) % mod
	}
	comb = func(n, m int) int {
		if m > n-m {
			return comb(n, n-m)
		}
		var num = 1
		var den = 1
		for i := 0; i < m; i += 1 {
			num *= n - i
			den *= i + 1
			num %= mod
			den %= mod
		}
		return (num * fastPower(den, mod-2)) % mod
	}

	m, n := abs(x2-x1), abs(y2-y1)
	if max(x1, x2) <= min(y1, y2) {
		fmt.Fprintln(w, comb(m+n, n))
		return
	}
	u := max(x1, x2) - min(y1, y2)
	fmt.Fprintln(w, (comb(m+n, n)-comb(m+n, u-1)+mod)%mod)
}

func LCDiDi2019012(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	// hilbert 曲线
	var hilbert func(res [][]int, n int, index int)
	hilbert = func(res [][]int, n int, index int) {
		if n == 1 {
			res[0][0] = index
			return
		}
		m := n * n / 4
		hilbert(res, n/2, index+m)
		for i := 0; i < n/2; i += 1 {
			for j := n / 2; j < n; j += 1 {
				res[i][j] = res[i][j-n/2] + m
			}
		}
		for i := n / 2; i < n; i += 1 {
			for j := 0; j < n/2; j += 1 {
				res[i][j] = 2*index - 1 + m*2 - res[n/2-1-j][i-n/2]
			}
			for j := n / 2; j < n; j += 1 {
				res[i][j] = 2*index + 4*m - 1 - res[i][n-1-j]
			}
		}
	}

	var n int
	fmt.Fscan(r, &n)
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	hilbert(res, n, 1)
	for _, row := range res {
		for _, e := range row {
			fmt.Fprint(w, e, " ")
		}
		fmt.Fprintln(w)
	}
}

func LCDiDi2020001(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	price := make([][]int, n)
	for i := range price {
		price[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(r, &price[i][j])
		}
	}
	ans := 0

	for i := 0; i < m; i++ {
		tmp := price[0][i]
		for j := 0; j < n; j++ {
			tmp = max(tmp, price[j][i])
		}
		ans += tmp
	}

	fmt.Fprintln(w, ans)

}

func LCDiDi2020002(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	nums, inp := make([]int, n), make([]string, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(r, &nums[i], &inp[i+1])
	}
	inp[0] = "+"
	fmt.Fscan(r, &nums[n-1])

	left, right := 0, 0
	for right < n {
		for right < n && inp[right] == inp[left] {
			right++
		}
		right--
		if inp[left] == "+" || inp[left] == "-" {
			if right < n-1 && (inp[right+1] == "*" || inp[right+1] == "/") {
				sort.Ints(nums[left:right])
			} else {
				sort.Ints(nums[left : right+1])
			}
		} else if inp[left] == "*" {
			if left > 0 && (inp[left-1] == "+" || inp[left-1] == "-") {
				sort.Ints(nums[left-1 : right+1])
			} else {
				sort.Ints(nums[left : right+1])
			}
		} else if inp[left] == "/" {
			sort.Ints(nums[left : right+1])
		}
		right++
		left = right
	}

	fmt.Fprint(w, nums[0], " ")
	for i := 1; i < n; i++ {
		fmt.Fprint(w, inp[i], " ", nums[i], " ")
	}
}

func LCDiDi2020003(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	arr := make([]byte, n)
	fmt.Fscan(r, &arr)
	ans := math.MaxInt64
	for i := 0; i <= 26; i++ {
		val, res := byte(i+'a'), 0
		for j := 0; j < n; {
			if arr[j] != val {
				j += m
				res++
				continue
			}
			j++
		}
		ans = min(ans, res)
	}
	fmt.Fprintln(w, ans)

}

func LCDiDi2020006(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(r, &arr[i])
	}
	ans := math.MaxInt64
	last, mmin := 0, 0
	for _, v := range arr[:m] {
		last += v
	}
	mmin = last
	for i := m; i < n; i++ {
		last += arr[i] - arr[i-m]
		ans = min(ans, mmin)
		mmin = min(mmin+arr[i], last)
	}
	fmt.Fprintln(w, min(ans, mmin))
}
