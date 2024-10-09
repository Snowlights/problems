package _00

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

func CF808A(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var n string
	fmt.Fscan(_r, &n)
	s0 := int64(n[0]-'0') + 1
	newS := strconv.FormatInt(s0, 10) + strings.Repeat("0", len(n)-1)
	new, _ := strconv.ParseInt(newS, 10, 64)
	old, _ := strconv.ParseInt(n, 10, 64)
	fmt.Fprintln(_w, new-old)
}

func CF808B(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var n, k int64
	fmt.Fscan(_r, &n, &k)
	arr := make([]int64, n)
	prefix := make([]int64, n+1)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(_r, &arr[i])
		prefix[i+1] = prefix[i] + arr[i]
	}

	weeks := n - k + 1
	total := int64(0)
	for i := int64(0); i < weeks; i++ {
		total += prefix[i+k] - prefix[i]
	}
	fmt.Fprintln(_w, total/weeks)
}

func CF808C(in io.Reader, out io.Writer) {
	_r := bufio.NewReader(in)
	_w := bufio.NewWriter(out)
	defer _w.Flush()

	var n, w int
	fmt.Fscan(_r, &n, &w)
	arr := make([]int, n)
	total := int(0)

	type cup struct {
		i   int
		cap int
		put int
	}
	cupList, cupMap := make([]*cup, 0, n), make(map[int]*cup)
	for i := 0; i < n; i++ {
		fmt.Fscan(_r, &arr[i])
		v := 0
		if arr[i]%2 == 1 {
			v = arr[i]/2 + 1
		} else {
			v = arr[i] / 2
		}
		total += v
		c := &cup{
			i:   i,
			cap: arr[i],
			put: v,
		}
		cupList = append(cupList, c)
		cupMap[c.i] = c
	}

	if total > w {
		fmt.Fprintln(_w, "-1")
		return
	}
	sort.Slice(cupList, func(i, j int) bool {
		return cupList[i].cap > cupList[j].cap
	})

	res := make([]int64, n)

	left := w - total
	for left > 0 {
		for i := 0; i < n; i++ {
			l := min(left, cupList[i].cap-cupList[i].put)
			left -= l
			cupList[i].put += l
			if left == 0 {
				break
			}
		}
	}

	for _, v := range cupList {
		res[v.i] = int64(v.put)
	}

	for _, v := range res {
		fmt.Fprint(_w, v, " ")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CF808D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	a := make([]int64, n)
	s := make([]int64, n+1)
	for i := range a {
		fmt.Fscan(r, &a[i])
		s[i+1] = s[i] + a[i]
	}
	if s[n]&1 > 0 {
		fmt.Fprint(w, "NO")
		return
	}
	mp := map[int64]bool{0: true}
	for i, v := range a {
		mp[v] = true
		if mp[s[i+1]-s[n]/2] {
			fmt.Fprint(w, "YES")
			return
		}
	}
	mp = map[int64]bool{0: true}
	for i := n - 1; i >= 0; i-- {
		mp[a[i]] = true
		if mp[s[n]/2-s[i]] {
			fmt.Fprint(w, "YES")
			return
		}
	}
	fmt.Fprint(w, "NO")
}

// dp[i][j] = max(dp[i][j], dp[i][j - weight[i]] + value[i]);

// dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);

const (
	weightOne   = 1
	weightTwo   = 2
	weightThree = 3
)

func CF808E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, b int
	fmt.Fscan(r, &n, &b)

	weightToValues := make(map[int][]int)
	weight, value := make([]int, n), make([]int, n)
	for i := range weight {
		fmt.Fscan(r, &weight[i], &value[i])
		weightToValues[weight[i]] = append(weightToValues[weight[i]], value[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(weightToValues[weightOne])))
	sort.Sort(sort.Reverse(sort.IntSlice(weightToValues[weightTwo])))
	sort.Sort(sort.Reverse(sort.IntSlice(weightToValues[weightThree])))

	one, two, three := weightToValues[weightOne], weightToValues[weightTwo], weightToValues[weightThree]

	dp := make([][]int, b+1)
	dp[0] = make([]int, 3)
	for i := 1; i <= b; i++ {
		dp[i] = make([]int, 3)
		dp[i][0] = math.MinInt64
		if dp[i-1][1] < len(one) && dp[i-1][0]+one[dp[i-1][1]] > dp[i][0] {
			dp[i][0] = dp[i-1][0] + one[dp[i-1][1]]
			dp[i][1] = dp[i-1][1] + 1
			dp[i][2] = dp[i-1][2]
		}
		if i > 1 && len(one)-dp[i-2][1] >= 2 && dp[i-2][0]+one[dp[i-2][1]]+one[dp[i-2][1]+1] > dp[i][0] {
			dp[i][0] = dp[i-2][0] + one[dp[i-2][1]] + one[dp[i-2][1]+1]
			dp[i][1] = dp[i-2][1] + 2
			dp[i][2] = dp[i-2][2]
		}
		if i > 1 && dp[i-2][2] < len(two) && dp[i-2][0]+two[dp[i-2][2]] > dp[i][0] {
			dp[i][0] = dp[i-2][0] + two[dp[i-2][2]]
			dp[i][1] = dp[i-2][1]
			dp[i][2] = dp[i-2][2] + 1
		}
	}
	ans := int(0)
	for i := 2; i <= b; i++ {
		dp[i][0] = max(dp[i][0], dp[i-1][0])
	}
	c := int(0)
	for i := (0); 3*i <= (b) && i <= len(three); i++ {
		if i > 0 {
			c += three[i-1]
		}
		ans = max(ans, int(c+dp[b-3*i][0]))
	}

	fmt.Fprintln(w, ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
