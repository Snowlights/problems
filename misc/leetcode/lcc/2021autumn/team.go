package _021autumn

// 1
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var m map[int]int

func numColor(root *TreeNode) int {
	m = make(map[int]int)
	dfs(root)
	return len(m)
}

func dfs(root *TreeNode) {
	if root != nil {
		m[root.Val] = 1
		dfs(root.Left)
		dfs(root.Right)
	}
}

// 2
func bicycleYard(position []int, terrain [][]int, obstacle [][]int) (ans [][]int) {

	//  骑行过程中速度不能为零或负值

	n, m := len(terrain), len(terrain[0])

	ok := make([][]bool, n)
	for i := range ok {
		ok[i] = make([]bool, m)
	}

	vis := make([][][110]bool, n)
	for i := range vis {
		vis[i] = make([][110]bool, m)
	}
	sx, sy := position[0], position[1]
	vis[sx][sy][1] = true
	type pair struct{ x, y, s int }
	q := []pair{{sx, sy, 1}}
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y, s := p.x, p.y, p.s
		if s == 1 {
			ok[x][y] = true
		}
		for _, d := range dir4 {
			if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m {
				ns := s + terrain[x][y] - terrain[xx][yy] - obstacle[xx][yy]
				if ns > 0 && !vis[xx][yy][ns] {
					vis[xx][yy][ns] = true
					q = append(q, pair{xx, yy, ns})
				}

			}
		}
	}

	ok[sx][sy] = false
	for i, r := range ok {
		for j, v := range r {
			if v {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return
}

// 3
type Node struct {
	a float64
	b int
}

func volunteerDeployment(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
	n := len(finalCnt) + 1
	arr := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	temp := make([]Node, 0)
	temp = append(temp, Node{
		a: 1,
		b: 0,
	})
	for i := 0; i < len(finalCnt); i++ {
		temp = append(temp, Node{a: 0, b: finalCnt[i]})
	}
	for i := len(plans) - 1; i >= 0; i-- {
		a, b := plans[i][0], plans[i][1]
		if a == 1 {
			temp[b] = Node{
				a: temp[b].a * 2,
				b: temp[b].b * 2,
			}
		} else if a == 2 {
			for j := 0; j < len(arr[b]); j++ {
				next := arr[b][j]
				temp[next] = Node{
					a: temp[next].a - temp[b].a,
					b: temp[next].b - temp[b].b,
				}
			}
		} else if a == 3 {
			for j := 0; j < len(arr[b]); j++ {
				next := arr[b][j]
				temp[next] = Node{
					a: temp[next].a + temp[b].a,
					b: temp[next].b + temp[b].b,
				}
			}
		}
	}
	a, b := float64(0), int64(0)
	for i := 0; i < len(temp); i++ {
		a = a + temp[i].a
		b = b + int64(temp[i].b)
	}
	per := float64(totalNum-b) / a
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = int(temp[i].a*per) + temp[i].b
	}
	return res
}

// 4
func securityCheck(a []int, k int) (ans int) {
	const mod int = 1e9 + 7
	dp := make([]int, k+1)
	dp[0] = 1
	c0 := 0
	for _, v := range a {
		v--
		if v == 0 {
			c0++
			continue
		}
		for s := k; s >= v; s-- {
			dp[s] = (dp[s] + dp[s-v]) % mod
		}
	}
	ans = dp[k]
	for i := 0; i < c0; i++ {
		ans = ans * 2 % mod
	}
	return
}

// 5
/*
黑一步胜 -> "Black"

黑一步不胜
- 白可以一步胜，不止一处位置 -> "White"
- 白仅有一处位置
-- 黑第一手下该位置 -> 白无法胜
- 白无获胜位置

白无法胜
- 白仅有一处位置时，黑第一手下白胜的位置
- 白无胜利位置时，枚举黑第一手和第二手位置 O(n)
- 第二手不止一处 -> "Black"
- 否则 "None"

*/
type pair struct{ x, y int }

var dir8 = []pair{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

const (
	black = "Black"
	white = "White"
	none  = "None"
)

func gobang(ps [][]int) (ans string) {
	color := make(map[pair]int, len(ps)+5)
	for _, p := range ps {
		p[2]++
		color[pair{p[0], p[1]}] = p[2]
	}
	// 1 黑
	// 2 白
	vis0 := map[pair]bool{}
	check := func(x, y, c int) bool {
		if vis0[pair{x, y}] {
			return false
		}
		vis0[pair{x, y}] = true
		for k, d := range dir8[:4] {
			cnt := 1
			// 检查一个方向
			for x, y := x+d.x, y+d.y; color[pair{x, y}] == c; x, y = x+d.x, y+d.y {
				cnt++
			}
			// 检查相反的另一方向
			d = dir8[k+4]
			for x, y := x+d.x, y+d.y; color[pair{x, y}] == c; x, y = x+d.x, y+d.y {
				cnt++
			}
			if cnt >= 5 {
				return true
			}
		}
		return false
	}

	// 黑一步胜
	for _, p := range ps {
		if p[2] == 2 {
			continue
		}
		x, y := p[0], p[1]
		for _, d := range dir8 {
			xx, yy := x+d.x, y+d.y
			if color[pair{xx, yy}] == 0 {
				if check(xx, yy, 1) {
					return black
				}
			}
		}
	}

	// 白能否一步胜
	vis0 = map[pair]bool{}
	setW := map[pair]bool{}
	pw := pair{}
	for _, p := range ps {
		if p[2] == 1 {
			continue
		}
		x, y := p[0], p[1]
		for _, d := range dir8 {
			xx, yy := x+d.x, y+d.y
			q := pair{xx, yy}
			if color[q] == 0 {
				if check(xx, yy, 2) {
					if setW[q] = true; len(setW) > 1 { // 不止一处位置，白可以一步胜
						return white
					}
					pw = q
				}
			}
		}
	}

	// 黑第一手下该位置 -> 白无法胜
	if len(setW) == 1 {
		color[pw] = 1 // 黑第一手下该位置
		vis0 = map[pair]bool{}
		setB := map[pair]bool{}

		x, y := pw.x, pw.y
		for _, d := range dir8 {
			xx, yy := x+d.x, y+d.y
			if color[pair{xx, yy}] == 0 {
				if check(xx, yy, 1) {
					if setB[pair{xx, yy}] = true; len(setB) > 1 {
						return black
					}
				}
			}
		}

		for _, p := range ps {
			if p[2] == 2 {
				continue
			}
			x, y := p[0], p[1]
			for _, d := range dir8 {
				xx, yy := x+d.x, y+d.y
				if color[pair{xx, yy}] == 0 {
					if check(xx, yy, 1) {
						if setB[pair{xx, yy}] = true; len(setB) > 1 {
							return black
						}
					}
				}
			}
		}
		return none
	}

	// 白防守
	// 黑第二手能否不止一处
	vis := map[pair]bool{}
	check3 := func(x, y int) bool {
		p0 := pair{x, y}
		if vis[p0] {
			return false
		}
		vis[p0] = true

		setB := map[pair]bool{}
		for di := range dir8 {
			xx, yy := x, y
			for k := 1; k < 6; k++ {
				d := dir8[di]
				xx += d.x
				yy += d.y
				if color[pair{xx, yy}] != 0 {
					continue
				}
				//if xx == 4 && yy==6 && x == 3 && y==7&&di==3 {
				//	println()
				//}
				cnt := 1
				// 检查一个方向
				for x, y := xx+d.x, yy+d.y; color[pair{x, y}] == 1; x, y = x+d.x, y+d.y {
					cnt++
				}
				// 检查相反的另一方向
				kk := di + 4
				if di >= 4 {
					kk = di - 4
				}
				d = dir8[kk]
				for x, y := xx+d.x, yy+d.y; color[pair{x, y}] == 1; x, y = x+d.x, y+d.y {
					cnt++
				}
				if cnt >= 5 {
					if setB[pair{xx, yy}] = true; len(setB) > 1 {
						return true
					}
				}
			}
		}
		return false
	}

	for _, p := range ps {
		if p[2] == 2 {
			continue
		}
		x, y := p[0], p[1]
		for i := -2; i <= 2; i++ {
			for j := -2; j <= 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				xx, yy := x+i, y+j
				q := pair{xx, yy}
				if color[q] != 0 {
					continue
				}
				color[q] = 1 // 黑落子
				if check3(xx, yy) {
					return black
				}
				delete(color, q)
			}
		}
	}

	return none
}

// 6
func ringGame(c []int64) int64 {
	ans := int64(1<<60 - 1)
	n := len(c)
	type pair struct {
		k, v int64
	}

	maxIndex := func() int {
		index, mx := 0, c[0]
		for i := 0; i < n; i++ {
			if c[i] > mx {
				mx = c[i]
				index = i
			}
		}
		return index
	}
	mx := maxIndex()
	for i := 59; i >= 0; i-- {
		pans := ans ^ (1 << i)
		tp := make([]*pair, 0, 10)
		for j := 0; j < n; j++ {
			cur, v := c[(j+mx)%n], 0
			if cur <= pans {
				cur |= pans
				v = 1
			}

			for len(tp) > 0 {
				if tp[len(tp)-1].v == 1 {
					if v == 1 {
						cur |= tp[len(tp)-1].k
						v = 1
						tp = tp[:len(tp)-1]
					} else if tp[len(tp)-1].k >= cur {
						cur |= tp[len(tp)-1].k
						v = 1
						tp = tp[:len(tp)-1]
					} else {
						break
					}
				} else if cur >= tp[len(tp)-1].k && v == 1 {
					cur |= tp[len(tp)-1].k
					v = 1
					tp = tp[:len(tp)-1]
				} else {
					break
				}
			}
			tp = append(tp, &pair{
				k: cur,
				v: int64(v),
			})
		}
		if len(tp) == 1 && tp[len(tp)-1].v == 1 {
			ans = pans
		}
	}

	return ans
}
