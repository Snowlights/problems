package main

import "math"

func lightSticks(height int, width int, indices []int) (ans []int) {

	// 火柴编码
	t, total, oneLine := math.MaxUint32, 2*height*width+height+width, 2*width+1
	point := func(x, y int) int {
		return (width+1)*x + y
	}
	top := func(x, y int) int {
		return (x-1)*oneLine + width + y
	}
	left := func(x, y int) int {
		return x*oneLine + y - 1
	}
	right := func(x, y int) int {
		return x*oneLine + y
	}
	down := func(x, y int) int {
		return x*oneLine + width + y
	}
	del := make([]bool, total)
	for _, v := range indices {
		del[v] = true
	}

	for i := 0; i <= height; i++ {
		for j := 0; j <= width; j++ {
			vis := append([]bool{}, del...)
			step, l := 0, total-len(indices)
			q := [][]int{{i, j}}

			for ; len(q) > 0; step++ {
				tmp := q
				q = nil

				for _, v := range tmp {
					x, y := v[0], v[1]
					if id := top(x, y); x > 0 && !vis[id] {
						vis[id] = true
						l--
						q = append(q, []int{x - 1, y})
					}
					if id := down(x, y); x < height && !vis[id] {
						vis[id] = true
						l--
						q = append(q, []int{x + 1, y})
					}
					if id := left(x, y); y > 0 && !vis[id] {
						vis[id] = true
						l--
						q = append(q, []int{x, y - 1})
					}
					if id := right(x, y); y < width && !vis[id] {
						vis[id] = true
						l--
						q = append(q, []int{x, y + 1})
					}
				}
			}
			if l > 0 {
				continue
			}

			if step < t {
				t = step
				ans = []int{point(i, j)}
			} else if step == t {
				ans = append(ans, point(i, j))
			}
		}
	}

	return
}
