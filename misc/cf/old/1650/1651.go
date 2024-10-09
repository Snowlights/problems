package _650

import (
	"bufio"
	"fmt"
	"io"
)

// https://codeforces.com/problemset/problem/1651/D
// https://codeforces.com/contest/1651/my
//输入 n(≤2e5) 和 n 个二维平面上的互不相同的整点，坐标范围 [1,2e5]。
//对每个整点，输出离它曼哈顿距离最近的，且不在输入中的整点。
//两点的曼哈顿距离=横坐标之差的绝对值+纵坐标之差的绝对值。

func CF1651D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	type pair struct {
		x, y int
	}
	type pair2 struct {
		x, y int
		pair
	}
	var n, x, y int
	fmt.Fscan(r, &n)
	pairList := make(map[pair]int, 0)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &x, &y)
		pairList[pair{
			x: x,
			y: y,
		}] = i
	}

	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	ans := make([]pair, n)
	q := []pair2{}

	for v, i := range pairList {
		for _, d := range dir {
			x, y := v.x+d[0], v.y+d[1]
			p := pair{x: x, y: y}
			if _, ok := pairList[p]; !ok {
				ans[i-1] = p
				q = append(q, pair2{x: v.x, y: v.y, pair: p})
				break
			}
		}
	}
	for _, p := range q {
		delete(pairList, pair{
			x: p.x,
			y: p.y,
		})
	}

	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, d := range dir {
				x, y := v.x+d[0], v.y+d[1]
				p := pair{
					x: x,
					y: y,
				}
				if i, ok := pairList[p]; ok {
					delete(pairList, p)
					ans[i-1] = v.pair
					q = append(q, pair2{
						x:    x,
						y:    y,
						pair: v.pair,
					})
				}
			}
		}
	}

	for _, v := range ans {
		fmt.Fprintln(w, v.x, v.y)
	}

}

func CF1651DV2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var n, x, y int
	fmt.Fscan(in, &n)
	s := make(map[pair]int, n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x, &y)
		s[pair{x, y}] = i
	}

	ans := make([]pair, n)
	type p2 struct {
		i, j int
		pair
	}
	q := []p2{}
	for v, i := range s {
		for _, d := range dir4 {
			w := pair{v.x + d.x, v.y + d.y}
			if s[w] == 0 {
				ans[i-1] = w
				q = append(q, p2{v.x, v.y, w})
				break
			}
		}
	}
	for _, p := range q {
		delete(s, pair{p.i, p.j})
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, d := range dir4 {
			w := pair{v.i + d.x, v.j + d.y}
			if i := s[w]; i > 0 {
				delete(s, w)
				ans[i-1] = v.pair
				q = append(q, p2{w.x, w.y, v.pair})
			}
		}
	}
	for _, p := range ans {
		fmt.Fprintln(out, p.x, p.y)
	}
}
