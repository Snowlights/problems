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

func CF1946C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ans := sort.Search(n/(k+1), func(low int) bool {
			low++
			cnt := 0
			var dfs func(int, int) int
			dfs = func(v, fa int) int {
				s := 1
				for _, w := range g[v] {
					if w != fa {
						s += dfs(w, v)
					}
				}
				if s >= low {
					cnt++
					return 0
				}
				return s
			}
			if dfs(0, -1) >= low {
				cnt++
			}
			return cnt <= k
		})
		Fprintln(out, ans)
	}

}

func main() { CF1946C(os.Stdin, os.Stdout) }
