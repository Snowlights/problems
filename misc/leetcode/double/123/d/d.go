package main

import (
	"math"
	"sort"
)

func numberOfPairs(points [][]int) (ans int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] > points[j][1]
	})
	for i, p := range points {
		y0 := p[1]
		maxY := math.MinInt
		for _, q := range points[i+1:] {
			y := q[1]
			if y <= y0 && y > maxY {
				maxY = y
				ans++
			}
		}
	}
	return
}
