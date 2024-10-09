package contest_260

import (
	"bufio"
	"fmt"
	"io"
)

func AtCoder260E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, a, b int64
	fmt.Fscan(r, &n, &m)
	h := make(map[int64][]int64)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &a, &b)
		h[a] = append(h[a], i)
		h[b] = append(h[b], i)
	}
	vis := make([]int64, m+1)
	j, cnt := int64(1), int64(0)
	ans := make([]int64, m+2)
	for i := int64(1); i <= m; i++ {
		for j <= m && cnt < n {
			for _, v := range h[j] {
				if vis[v] == 0 {
					cnt++
				}
				vis[v]++
			}
			j++
		}
		if cnt < n {
			break
		}
		ans[j-i]++
		ans[m-i+2]--
		for _, v := range h[i] {
			vis[v]--
			if vis[v] == 0 {
				cnt--
			}
		}
	}
	for i := int64(1); i <= m; i++ {
		ans[i] += ans[i-1]
		fmt.Fprint(w, ans[i], " ")
	}

}
