package main

import "sort"

func minTime(skill, mana []int) int64 {
	n := len(skill)
	lastFinish := make([]int, n) // 第 i 名巫师完成上一瓶药水的时间
	for _, m := range mana {
		// 按题意模拟
		sumT := 0
		for i, x := range skill {
			sumT = max(sumT, lastFinish[i]) + x*m
		}
		// 倒推：如果酿造药水的过程中没有停顿，那么 lastFinish[i] 应该是多少
		lastFinish[n-1] = sumT
		for i := n - 2; i >= 0; i-- {
			lastFinish[i] = lastFinish[i+1] - skill[i+1]*m
		}
	}
	return int64(lastFinish[n-1])
}

func minTime2(skill, mana []int) int64 {
	n, m := len(skill), len(mana)
	s := make([]int, n+1) // skill 的前缀和
	for i, x := range skill {
		s[i+1] = s[i] + x
	}

	start := 0
	for j := 1; j < m; j++ {
		mx := 0
		for i := range n {
			mx = max(mx, mana[j-1]*s[i+1]-mana[j]*s[i])
		}
		start += mx
	}
	return int64(start + mana[m-1]*s[n])
}

func minTime3(skill, mana []int) int64 {
	n, m := len(skill), len(mana)
	s := make([]int, n+1)
	for i, x := range skill {
		s[i+1] = s[i] + x
	}

	suf := []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		if skill[i] > skill[suf[len(suf)-1]] {
			suf = append(suf, i)
		}
	}

	pre := []int{0}
	for i := 1; i < n; i++ {
		if skill[i] > skill[pre[len(pre)-1]] {
			pre = append(pre, i)
		}
	}

	start := 0
	for j := 1; j < m; j++ {
		record := suf
		if mana[j-1] < mana[j] {
			record = pre
		}
		mx := 0
		for _, i := range record {
			mx = max(mx, mana[j-1]*s[i+1]-mana[j]*s[i])
		}
		start += mx
	}
	return int64(start + mana[m-1]*s[n])
}

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }

// 计算 points 的上凸包
// 由于横坐标是严格递增的，无需排序
func convexHull(points []vec) (q []vec) {
	for i := len(points) - 1; i >= 0; i-- {
		p := points[i]
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
	return
}

func minTime4(skill, mana []int) int64 {
	n, m := len(skill), len(mana)
	s := make([]int, n+1)
	ps := make([]vec, n)
	for i, x := range skill {
		s[i+1] = s[i] + x
		ps[i] = vec{s[i], x}
	}
	ps = convexHull(ps) // 去掉无用数据

	start := 0
	for j := 1; j < m; j++ {
		p := vec{mana[j-1] - mana[j], mana[j-1]}
		// p.dot(ps[i]) 是个上凸函数，二分找最大值
		i := sort.Search(len(ps)-1, func(m int) bool { return p.dot(ps[m]) > p.dot(ps[m+1]) })
		start += p.dot(ps[i])
	}
	return int64(start + mana[m-1]*s[n])
}
