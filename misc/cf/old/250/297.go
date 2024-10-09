package _50

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

//	由于偶串，不管怎样变换，最多不会超出原来1的个数，
//
// 而奇串最多不会超出原来1的个数加1.而任意字串的1的个数小于A最多1的个数，都可以通过变换得到！
func CF297A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var a, b string
	fmt.Fscan(r, &a, &b)
	if (strings.Count(a, "1")+1)/2*2 >= strings.Count(b, "1") {
		fmt.Fprintln(w, "YES")
		return
	}

	fmt.Fprintln(w, "NO")
}

func CF297B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, k int64
	fmt.Fscan(r, &n, &m, &k)
	if n > m {
		fmt.Fprintln(out, "YES")
		return
	}
	a, b := make([]int64, n), make([]int64, m)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	for i := int64(0); i < m; i++ {
		fmt.Fscan(r, &b[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	for i, v := range a {
		if v > b[i] {
			fmt.Fprintln(out, "YES")
			return
		}
	}

	fmt.Fprintln(out, "NO")
}
