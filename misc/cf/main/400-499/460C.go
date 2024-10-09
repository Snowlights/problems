//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// 一个长度为 n 的序列 a ，你有 m 次操作的机会，
// 每次操作是将其中连续的 w 个元素增加 1 。最大化最终序列的最小值。
func CF460C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, w int
	Fscan(in, &n, &m, &w)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	Fprintln(out, sort.Search(1e9+m, func(low int) bool {
		low++                    // 改为二分最小的不满足要求的值，这样 sort.Search 返回的就是最大的满足要求的值
		diff := make([]int, n+1) // 差分数组
		sumD, need := 0, 0       // 总的差分值 修改的次数
		for i, v := range a {
			sumD += diff[i]
			if sumD+v < low {
				d := low - sumD - v
				need += d
				if need > m {
					return true
				}
				sumD += d
				diff[min(i+w, n)] -= d
			}
		}
		return false
	}))

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() { CF460C(os.Stdin, os.Stdout) }
