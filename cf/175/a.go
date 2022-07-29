package _75

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func CF175C(in io.Reader, out io.Writer) {

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	type master struct {
		num, score int64
	}
	var N int64
	fmt.Fscan(r, &N)
	masterList := make([]*master, 0, N)
	for i := int64(0); i < N; i++ {
		var n, s int64
		fmt.Fscan(r, &n, &s)
		masterList = append(masterList, &master{
			num:   n,
			score: s,
		})
	}
	var T int64
	fmt.Fscan(r, &T)
	arr := make([]int64, T)
	for i := int64(0); i < T; i++ {
		fmt.Fscan(r, &arr[i])
	}
	newArr := make([]int64, T+1)
	newArr[0] = arr[0]
	for i := int64(1); i < T; i++ {
		newArr[i] = arr[i] - arr[i-1]
	}
	newArr[T] = 1e12

	sort.Slice(masterList, func(i, j int) bool {
		return masterList[i].score < masterList[j].score
	})

	ans, idx := int64(0), int64(0)
	for _, m := range masterList {
		for m.num > 0 && idx < T+1 {
			v := min(m.num, newArr[idx])
			ans += v * m.score * (idx + 1)
			newArr[idx] -= v
			if newArr[idx] == 0 {
				idx++
			}
			m.num -= v
		}
		if idx == T+1 {
			break
		}
	}
	fmt.Fprintln(w, ans)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
