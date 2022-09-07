package autoX

// 4


// 点
type Point struct {
	x, y int
}

// 圆
type Circle struct {
	x, y, r int
}

func judge(p1, p2 *Point, c *Circle) bool {
	flag1 := (p1.x-c.x)*(p1.x-c.x)+(p1.y-c.y)*(p1.y-c.y) <= c.r*c.r
	flag2 := (p2.x-c.x)*(p2.x-c.x)+(p2.y-c.y)*(p2.y-c.y) <= c.r*c.r
	//情况一、两点都在圆内 :一定不相交
	if flag1 && flag2 {
		return false
	} else if flag1 || flag2 {
		//情况二、一个点在圆内，一个点在圆外：一定相交
		return true
	} else {
		//将直线p1p2化为一般式：Ax+By+C=0的形式。先化为两点式，然后由两点式得出一般式
		A:=p1.y-p2.y
		B:=p2.x-p1.x
		C:=p1.x*p2.y-p2.x*p1.y
		//使用距离公式判断圆心到直线ax+by+c=0的距离是否大于半径
		dist1:=A*c.x+B*c.y+C
		dist1*=dist1
		dist2:=(A*A+B*B)*c.r*c.r
		if dist1 > dist2 {
			return false
		}
		//点到直线距离大于半径r  不相交
		angle1:=(c.x-p1.x)*(p2.x-p1.x)+(c.y-p1.y)*(p2.y-p1.y)
		angle2:=(c.x-p2.x)*(p1.x-p2.x)+(c.y-p2.y)*(p1.y-p2.y)
		if angle1 > 0 && angle2 > 0 {
			return true
		}
		//余弦都为正，则是锐角 相交
		return false
	}
}

type Line struct {
	x1, y1, x2, y2 int
}

func intersection(l1, l2 *Line) bool {
	//快速排斥实验
	if (max(l1.x1, l1.x2) < (min(l2.x1, l2.x2))) ||
		(max(l1.y1, l1.y2) < (min(l2.y1, l2.y2))) ||
		(max(l2.x1, l2.x2) < (min(l1.x1, l1.x2))) ||
		(max(l2.y1, l2.y2) < (min(l1.y1, l1.y2))){
		return false
	}
	//跨立实验
	if (((l1.x1 - l2.x1)*(l2.y2 - l2.y1) - (l1.y1 - l2.y1)*(l2.x2 - l2.x1))*
		((l1.x2 - l2.x1)*(l2.y2 - l2.y1) - (l1.y2 - l2.y1)*(l2.x2 - l2.x1))) > 0 ||
		(((l2.x1 - l1.x1)*(l1.y2 - l1.y1) - (l2.y1 - l1.y1)*(l1.x2 - l1.x1))*
			((l2.x2 - l1.x1)*(l1.y2 - l1.y1) - (l2.y2 - l1.y1)*(l1.x2 - l1.x1))) > 0 {
		return false
	}

	return true
}

// 两圆是否相交
func isIntersection(c1, c2 *Circle) bool {
	r := c1.r + c2.r
	x := c1.x - c2.x
	y := c1.y - c2.y
	d := x*x + y*y
	return d <= r * r
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

func cal(a, b []int) bool {
	switch len(a) {
	case 3:
		switch len(b) {
		case 3:
			return isIntersection(&Circle{
				x: int(a[0]),
				y: int(a[1]),
				r: int(a[2]),
			}, &Circle{
				x: int(b[0]),
				y: int(b[1]),
				r: int(b[2]),
			})
		case 4:
			return judge(&Point{
				x: int(b[0]),
				y: int(b[1]),
			}, &Point{
				x: int(b[2]),
				y: int(b[3]),
			}, &Circle{
				x: int(a[0]),
				y: int(a[1]),
				r: int(a[2]),
			})
		}
	case 4:
		switch len(b) {
		case 3:
			return judge(&Point{
				x: int(a[0]),
				y: int(a[1]),
			}, &Point{
				x: int(a[2]),
				y: int(a[3]),
			}, &Circle{
				x: int(b[0]),
				y: int(b[1]),
				r: int(b[2]),
			})
		case 4:
			return intersection(&Line{
				x1: int(a[0]),
				y1: int(a[1]),
				x2: int(a[2]),
				y2: int(a[3]),
			}, &Line{
				x1: int(b[0]),
				y1: int(b[1]),
				x2: int(b[2]),
				y2: int(b[3]),
			})
		}
	}
	return false
}

func antPass(geometry [][]int, path [][]int) []bool {
	g, vis := make(map[int][]int), make(map[int]bool)
	for i := 0; i < len(geometry); i++ {
		for j := i + 1; j < len(geometry); j++ {
			if cal(geometry[i], geometry[j]) {
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
	}

	gMap := make([]map[int]bool, 0)
	for k := range geometry {
		if vis[k] {
			continue
		}
		q, ans := []int{k}, make(map[int]bool)
		vis[k] = true
		ans[k] = true
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, vv := range g[v] {
					if ans[vv] {
						continue
					}
					vis[vv] = true
					ans[vv] = true
					q = append(q, vv)
				}
			}
		}
		gMap = append(gMap, ans)
	}

	res := make([]bool, 0, len(path))
	for _, p := range path {
		found := false
		for _, g := range gMap {
			if g[p[0]] && g[p[1]] {
				res = append(res, true)
				found = true
				break
			}
		}
		if !found {
			res = append(res, false)
		}
	}
	return res
}
