package template

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"strconv"
)

func LuoGuP3865(in io.Reader, out io.Writer) {
	// r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	r := bufio.NewScanner(in)
	r.Split(bufio.ScanWords)

	rInt := func() int {
		r.Scan()
		x, _ := strconv.Atoi(string(r.Bytes()))
		return x
	}

	n, q := rInt(), rInt()
	a := make([]int, n)
	for i := range a {
		a[i] = rInt()
	}
	st := NewST(a)
	for i := 0; i < q; i++ {
		left, right := rInt(), rInt()
		fmt.Fprintln(w, st.Query(left-1, right))
	}
}

type ST [][]int

func NewST(a []int) ST {
	n := len(a)
	sz := bits.Len(uint(n))
	st := make(ST, n)
	for i, v := range a {
		st[i] = make([]int, sz)
		st[i][0] = v
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = st.Op(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

// 查询区间 [l,r)，注意 l 和 r 是从 0 开始算的
func (st ST) Query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	return st.Op(st[l][k], st[r-1<<k][k])
}

// min, max, gcd, ...
func (ST) Op(a, b int) int {
	return max(a, b)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
