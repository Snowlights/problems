package main

import "sort"

func survivedRobotsHealths(positions []int, healths []int, directions string) (ans []int) {
	type pair struct {
		p, h, i int
		d       byte
	}
	const (
		Left  = 'L'
		Right = 'R'
	)

	pList := make([]*pair, 0, len(positions))
	for i, p := range positions {
		pList = append(pList, &pair{
			p: p,
			h: healths[i],
			i: i,
			d: directions[i],
		})
	}
	sort.Slice(pList, func(i, j int) bool {
		return pList[i].p < pList[j].p
	})

	var stack []*pair
	for _, p := range pList {
		if len(stack) == 0 || stack[len(stack)-1].d == p.d {
			stack = append(stack, p)
			continue
		}

		if p.d == Left {
			for len(stack) > 0 && stack[len(stack)-1].d == Right {
				if stack[len(stack)-1].h == p.h {
					stack = stack[:len(stack)-1]
					p.h = 0
					break
				}
				if stack[len(stack)-1].h > p.h {
					stack[len(stack)-1].h--
					p.h = 0
					break
				} else {
					stack = stack[:len(stack)-1]
					p.h--
					if p.h == 0 {
						break
					}
				}
			}
		}
		if p.h > 0 {
			stack = append(stack, p)
		}
	}

	sort.Slice(stack, func(i, j int) bool {
		return stack[i].i < stack[j].i
	})
	for _, v := range stack {
		ans = append(ans, v.h)
	}

	return
}
