package meituan

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// meituan-001
func LCMeiTuan001(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	var s string
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &s)

		if len(s) == 0 {
			fmt.Fprintln(w, "Wrong")
			continue
		}

		if (s[0] >= 'a' && s[0] <= 'z') || (s[0] >= 'A' && s[0] <= 'Z') {
			shuzi, all := false, true
			for i := 1; i < len(s); i++ {
				if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
					if s[i] >= '0' && s[i] <= '9' {
						shuzi = true
					}
				} else {
					all = false
					break
				}
			}
			if shuzi && all {
				fmt.Fprintln(w, "Accept")
			} else {
				fmt.Fprintln(w, "Wrong")
			}
		} else {
			fmt.Fprintln(w, "Wrong")
		}

	}

}

func LCMeiTuan002(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	ww := make([]int, n)
	for i := range ww {
		fmt.Fscan(r, &ww[i])
	}
	q := make([]int, n)
	for i := range q {
		fmt.Fscan(r, &q[i])
	}

	sum, f := make([]int, n+1), make([]int, n+1)
	for i := range f {
		f[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if f[x] != x {
			f[x] = find(f[x])
		}
		return f[x]
	}
	ans := make([]int, n)
	for i := n - 1; i > 0; i-- {
		x := q[i] - 1
		to := find(x + 1)
		f[x] = to
		sum[to] += sum[x] + ww[x]
		ans[i-1] = max(ans[i], sum[to])
	}

	for _, v := range ans {
		fmt.Fprintln(w, v)
	}
}

func LCMeiTuan003(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var m, n, v, p int
	fmt.Fscan(r, &n, &m)
	type pair struct {
		idx, count int
	}
	pList := make([]*pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &v, &p)
		pList[i] = &pair{
			idx:   i,
			count: v + p*2,
		}
	}

	sort.Slice(pList, func(i, j int) bool {
		return pList[i].count > pList[j].count || (pList[i].count == pList[j].count &&
			pList[i].idx < pList[j].idx)
	})
	arr := pList[:m]
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].idx < arr[j].idx
	})

	for _, v := range arr {
		fmt.Fprint(w, v.idx+1, " ")
	}
}

func LCMeiTuan004(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, op, k, x, y int
	fmt.Fscan(r, &n)
	A, B := make([]int, n), make([]int, n)
	for i := range A {
		fmt.Fscan(r, &A[i])
		B[i] = -1
	}

	for fmt.Fscan(r, &m); m > 0; m-- {
		fmt.Fscan(r, &op)
		switch op {
		case 1:
			fmt.Fscan(r, &k, &x, &y)
			copy(B[y-1:], A[x-1:x-1+k])
		case 2:
			fmt.Fscan(r, &x)
			fmt.Fprintln(w, B[x-1])
		}
	}

}

func LCMeiTuan005(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k, u, v int
	g := make(map[int][]int)
	fmt.Fscan(r, &n, &k)
	arr := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(r, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	node, level, mod := 0, 0, int(1e9+7)
	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(idx int) int {
		if 0 > idx || idx >= n {
			return 0
		}
		if arr[idx] < level || arr[idx] > level+k {
			return 0
		}
		if arr[idx] == level && idx < node {
			return 0
		}
		res := 1
		vis[idx] = true
		for _, next := range g[idx+1] {
			if !vis[next-1] {
				res *= (dfs(next-1) + 1)
				res %= mod
			}
		}
		vis[idx] = false
		return res
	}
	ans := 0
	for i := 0; i < n; i++ {
		node = i
		level = arr[i]
		vis[i] = true
		ans += dfs(i)
		ans %= mod
	}
	fmt.Println(ans)

}

func LCMeiTuan006(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	var T string
	fmt.Fscan(r, &n, &T)
	PM := strings.Index(T, "M")
	PT := strings.Index(T[PM:], "T") + PM
	ST := strings.LastIndex(T, "T")
	SM := strings.LastIndex(T[:ST], "M")
	fmt.Println(T[PT+1 : SM])
}

func LCMeiTuan007(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	ans, h := make([]int, 0, n), make(map[int]bool)
	for j := 0; j < n; j++ {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &arr[i])
		}
		for _, v := range arr {
			if h[v] {
				continue
			}
			ans = append(ans, v)
			h[v] = true
			break
		}
	}
	for _, v := range ans {
		fmt.Fprint(w, v, " ")
	}
}

func LCMeiTuan008(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, x, y, u, v int
	fmt.Fscan(r, &n, &x, &y)
	g := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(r, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	a, b := make([]int, n+1), make([]int, n+1)
	var dfs func(int, []int)
	dfs = func(from int, arr []int) {
		vis := make(map[int]bool)
		q, idx := []int{from}, 0
		vis[from] = true
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				arr[v] = idx
				for _, vv := range g[v] {
					if vis[vv] {
						continue
					}
					q = append(q, vv)
					vis[vv] = true
				}
			}
			idx++
		}
	}
	dfs(x, a)
	dfs(y, b)
	ans := 0
	for i := 1; i <= n; i++ {
		if a[i] > b[i] {
			ans = max(ans, a[i])
		}
	}
	fmt.Fprintln(w, ans)
}

func LCMeiTuan009(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	mod := 998244353
	var n, m int
	fmt.Fscan(r, &n, &m)
	// 长度为i时，总价为j时方案数
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k*j <= n; k++ {
				dp[i][j*k] = (dp[i][j*k] + dp[i-1][j]) % mod
			}
		}
	}

	ans := 0
	for _, v := range dp[m-1] {
		ans = (ans + v) % mod
	}
	fmt.Fprintln(w, ans)
}

func LCMeiTuan0010(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &nums[i])
	}

	var res int
	for l := 1; l <= m; l++ {
		r := sort.Search(m+1, func(x int) bool {
			if x < l {
				return false
			}
			var pre int
			for i := 0; i < n; i++ {
				if nums[i] < l || nums[i] > x {
					if nums[i] >= pre {
						pre = nums[i]
					} else {
						return false
					}
				}
			}
			return true
		})
		res += m - r + 1
	}
	fmt.Fprintln(w, res)

}

func LCMeiTuan0011(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var a, b, c, d, e, f, g int
	fmt.Fscan(r, &a, &b, &c, &d, &e, &f, &g)
	var res int
	var x int
	for d > 0 && (a > 0 || b > 0 || c > 0) {
		if a > 0 && e >= f && e >= g {
			x = min(a, d)
			res += e * x
			a -= x
			d -= x
			e = -1
		} else if b > 0 && f >= e && f >= g {
			x = min(b, d)
			res += f * x
			b -= x
			d -= x
			f = -1
		} else if c > 0 && g >= e && g >= f {
			x = min(c, d)
			res += g * x
			c -= x
			d -= x
			g = -1
		}
	}
	fmt.Println(res)
}

func LCMeiTuan0012(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var m, n, q, op, x, y int
	fmt.Fscan(r, &m, &n)
	bookstore := make(map[int]int)
	locked := make([]bool, n+1)
	tuan := make(map[int]bool)

	for fmt.Fscan(r, &q); q > 0; q-- {
		fmt.Fscan(r, &op)
		switch op {
		case 1:
			fmt.Fscan(r, &x, &y)
			if locked[y] || tuan[x] {
				continue
			}
			v, ok := bookstore[x]
			if ok && locked[v] {
				continue
			}
			bookstore[x] = y

		case 2:
			fmt.Fscan(r, &y)
			locked[y] = true
		case 3:
			fmt.Fscan(r, &y)
			locked[y] = false
		case 4:
			fmt.Fscan(r, &x)
			v, ok := bookstore[x]
			if ok && !locked[v] {
				delete(bookstore, x)
				tuan[x] = true
				fmt.Fprintln(w, v)
			} else {
				fmt.Fprintln(w, -1)
			}

		case 5:
			fmt.Fscan(r, &x)
			tuan[x] = false
		}
	}

}

func LCMeiTuan0013(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	var s string
	fmt.Fscan(r, &n, &s)
	ans, tmp := 0, 0
	for _, v := range s {
		if v == 'e' {
			tmp++
			ans = max(ans, tmp)
		} else {
			tmp--
			if tmp < 0 {
				tmp = 0
			}
		}
	}
	fmt.Fprintln(w, ans)
}

func LCMeiTuan0014(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var x, y int
	fmt.Fscan(r, &x, &y)
	if x == y {
		fmt.Fprintln(w, fmt.Sprintf("%s%s", strings.Repeat("A", x), strings.Repeat("B", y)))
		return
	}

	n := x + y
	nums1 := make([]int, n)
	nums2 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &nums1[i])
	}
	copy(nums2, nums1)
	sort.Ints(nums2)
	m1 := "A"
	m2 := "B"
	var midNum int
	if x > y {
		midNum = nums2[x]
		m1, m2 = m2, m1
	} else {
		midNum = nums2[y]
	}
	for _, num := range nums1 {
		if num >= midNum {
			fmt.Fprint(w, m1)
		} else {
			fmt.Fprint(w, m2)
		}
	}

}

func LCMeiTuan0015(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, xs, ys, xt, yt int
	fmt.Fscan(r, &n, &m, &xs, &ys, &xt, &yt)
	xs--
	ys--
	xt--
	yt--
	a, b, vis := make([][]int, n), make([][]int, n), make([][]bool, n)
	for i := range a {
		a[i] = make([]int, m)
		b[i] = make([]int, m)
		vis[i] = make([]bool, m)
	}
	for i := range a {
		for j := range a[i] {
			fmt.Fscan(r, &a[i][j])
		}
	}
	for i := range b {
		for j := range b[i] {
			fmt.Fscan(r, &b[i][j])
		}
	}
	q, time := [][]int{{xs, ys}}, 0
	vis[xs][ys] = true
	for len(q) > 0 {
		tmp := q
		for _, v := range tmp {
			x, y := v[0], v[1]
			if x == xt && y == yt {
				fmt.Fprintln(w, time)
				return
			}

			// 当时间处在 [k * a[i][j] + k * b[i][j], (k+1) * a[i][j] + k * b[i][j])时，
			// 行人可以选择走到 i±1 行 j 列的路口。
			if time%(a[x][y]+b[x][y]) < a[x][y] {
				if x > 0 && !vis[x-1][y] {
					vis[x-1][y] = true
					q = append(q, []int{x - 1, y})
				}
				if x < n-1 && !vis[x+1][y] {
					vis[x+1][y] = true
					q = append(q, []int{x + 1, y})
				}
			}
			// 当时间处在 [(k+1) * a[i][j] + k * b[i][j], (k+1) * a[i][j] + (k+1) * b[i][j])时，
			// 行人可以选择走到 i 行 j±1 列的路口。
			if time >= a[x][y] && (time-a[x][y])%(a[x][y]+b[x][y]) < b[x][y] {
				if y > 0 && !vis[x][y-1] {
					vis[x][y-1] = true
					q = append(q, []int{x, y - 1})
				}
				if y < m-1 && !vis[x][y+1] {
					vis[x][y+1] = true
					q = append(q, []int{x, y + 1})
				}
			}
		}
		time++
	}

}
