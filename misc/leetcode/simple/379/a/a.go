package main

func areaOfMaxDiagonal(dimensions [][]int) int {
	mx, ans := 0, 0
	for _, v := range dimensions {
		if t, area := v[0]*v[0]+v[1]*v[1], v[0]*v[1]; t > mx || t == mx && area > ans {
			ans = area
			mx = t
		}
	}
	return ans
}
