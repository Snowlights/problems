package simple_290

import "sort"

func countRectangles(rectangles [][]int, points [][]int) []int {
	sort.Slice(rectangles, func(i, j int) bool { return rectangles[i][0] > rectangles[j][0] })
	for i := range points {
		points[i] = append(points[i], i)
	}
	sort.Slice(points, func(i, j int) bool { return points[i][0] > points[j][0] })

	ans := make([]int, len(points))
	i, cnt := 0, [101]int{}
	for _, p := range points {
		for ; i < len(rectangles) && rectangles[i][0] >= p[0]; i++ {
			cnt[rectangles[i][1]]++
		}
		for _, c := range cnt[p[1]:] {
			ans[p[2]] += c
		}
	}
	return ans
}
