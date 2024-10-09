package yinglian

import "sort"

// https://leetcode.cn/contest/cnunionpay2022/

// 1
type ListNode struct {
	Val  int
	Next *ListNode
}

func reContruct(head *ListNode) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := newHead
	for cur.Next != nil {
		if cur.Next.Val%2 == 0 {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return newHead.Next
}

// 2
func explorationSupply(station []int, pos []int) []int {
	ans := make([]int, 0, len(station))

	for _, p := range pos {
		i := sort.SearchInts(station, p)
		j := -1
		if i > 0 {
			j = i - 1
		}
		if j < 0 || i < len(station) && station[i]-p < p-station[j] {
			j = i
		}
		ans = append(ans, j)
	}

	return ans
}

// 3
func storedEnergy(storeLimit int, power []int, supply [][]int) (ans int) {
	j := 0
	for i, s := range supply {
		t := len(power)
		if i+1 < len(supply) {
			t = supply[i+1][0]
		}
		for min, max := s[1], s[2]; j < t; j++ {
			if p := power[j]; p < min {
				if ans -= min - p; ans < 0 {
					ans = 0
				}
			} else if p > max {
				if ans += p - max; ans > storeLimit {
					ans = storeLimit
				}
			}
		}
	}
	return
}

// 4
type VendingMachine struct {
}

var (
	m   map[string][]pair
	dis map[string]int
)

func Constructor() (_ VendingMachine) {
	m = map[string][]pair{}
	dis = map[string]int{}
	return
}

func (VendingMachine) AddItem(time int, number int, item string, price int, duration int) {
	m[item] = append(m[item], pair{number, time + duration, price})
}

func (VendingMachine) Sell(time int, customer string, item string, number int) int64 {
	ans := 0

	ps := m[item]
	if ps == nil {
		return -1
	}

	c := 0
	for _, p := range ps {
		if p.end >= time {
			c += p.num
		}
	}
	if c < number {
		return -1
	}

	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.pr < b.pr || a.pr == b.pr && a.end < b.end })
	s := 0
	for number > 0 {
		p := ps[0]
		if p.end < time {
			ps = ps[1:]
			continue
		}
		if p.num > number {
			s += number * p.pr
			p.num -= number
			ps[0] = p
			number = 0
		} else {
			s += p.num * p.pr
			number -= p.num
			ps = ps[1:]
		}
	}

	m[item] = ps

	if dis[customer] == 0 {
		dis[customer] = 100
	}

	ans = (s*dis[customer] + 99) / 100

	dis[customer]--
	return int64(ans)
}

type pair struct{ num, end, pr int }
